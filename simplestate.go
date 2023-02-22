package simplestate

import (
	"fmt"
	"sync"
)

var (
	_ error = (*InvalidStateError[string])(nil)
	_ error = (*InvalidEventError[string])(nil)
	_ error = (*SameStateTransitionError[string])(nil)
)

type (
	// Machine is the simple state machine.
	Machine[Tstate comparable, Tevent comparable] struct {
		sync.RWMutex
		events       map[Tevent]map[Tstate]Tstate
		currentState Tstate
		settings     Settings
		callback     func(transition Transition[Tstate, Tevent])
	}

	// Settings is the settings for the state machine.
	Settings struct {
		// AllowSameStateTransition allows the state machine to transition to the same state.
		AllowSameStateTransition bool
	}

	// Event describes an event that can be triggered.
	// An event can have multiple sources and a single destination.
	// An event can be triggered if the current state is one of the sources.
	// When an event is triggered, the current state is changed to the destination.
	// State and event identifiers are generic types for better compile time checks.
	Event[Tstate comparable, Tevent comparable] struct {
		Event       Tevent
		Sources     []Tstate
		Destination Tstate
	}

	// Events is a list of events.
	Events[Tstate comparable, Tevent comparable] []Event[Tstate, Tevent]

	// Transition describes a transition from one state to another.
	// A transition is triggered by an event.
	// A transition has a single source and a destination.
	// State and event identifiers are generic types for better compile time checks.
	Transition[Tstate comparable, Tevent comparable] struct {
		Event       Tevent
		Source      Tstate
		Destination Tstate
	}

	// InvalidEventError is returned when an invalid event is triggered.
	InvalidEventError[Tevent comparable] struct {
		Event Tevent
	}

	// InvalidStateError is returned when an invalid state is set.
	InvalidStateError[Tstate comparable] struct {
		State Tstate
	}

	// SameStateTransitionError is returned when a transition to the same state is triggered.
	SameStateTransitionError[Tevent comparable] struct {
		Event Tevent
	}
)

// NewMachine creates a new state machine.
// The start state is the initial state of the state machine.
// The events are the events that can be triggered.
// The settings are the settings for the state machine.
// The callback is called when a transition is triggered.
func NewMachine[Tstate comparable, Tevent comparable](startState Tstate, events []Event[Tstate, Tevent], settings Settings, callback func(args Transition[Tstate, Tevent])) *Machine[Tstate, Tevent] {
	eventMap := make(map[Tevent]map[Tstate]Tstate)
	for _, e := range events {
		eventMap[e.Event] = make(map[Tstate]Tstate)
		for _, s := range e.Sources {
			eventMap[e.Event][s] = e.Destination
		}
	}

	return &Machine[Tstate, Tevent]{
		events:       eventMap,
		currentState: startState,
		settings:     settings,
		callback:     callback,
	}
}

// Get returns the current state.
func (m *Machine[Tstate, Tevent]) Get() Tstate {
	m.RLock()
	defer m.RUnlock()
	return m.currentState
}

// Is returns true if the current state is the given state.
func (m *Machine[Tstate, Tevent]) Is(state Tstate) bool {
	m.RLock()
	defer m.RUnlock()
	return m.currentState == state
}

// Set sets the current state without triggering any events.
func (m *Machine[Tstate, Tevent]) Set(state Tstate) {
	m.Lock()
	defer m.Unlock()
	m.currentState = state
}

// Can returns true if the given event can be triggered.
func (m *Machine[Tstate, Tevent]) Can(event Tevent) bool {
	m.RLock()
	defer m.RUnlock()
	if _, err := m.canTrigger(event); err != nil {
		return false
	}
	return true
}

// Trigger triggers the given event.
// If the event can be triggered, the current state is changed to the destination state.
// If the event cannot be triggered, an error is returned.
// The error can be one of the following:
//   - InvalidEventError: the event is not defined
//   - InvalidStateError: the current state is not a source state
//   - SameStateTransitionError: the current state is the same as the destination state
//     and the state machine is configured to not allow same state transitions
func (m *Machine[Tstate, Tevent]) Trigger(event Tevent) error {
	m.Lock()

	transition, err := m.canTrigger(event)
	if err != nil {
		m.Unlock()
		return err
	}

	m.currentState = transition.Destination
	m.Unlock()
	m.callback(*transition)
	return nil
}

func (m *Machine[Tstate, Tevent]) canTrigger(event Tevent) (*Transition[Tstate, Tevent], error) {
	destinationStates, ok := m.events[event]
	if !ok {
		return nil, &InvalidEventError[Tevent]{Event: event}
	}

	destinationState, ok := destinationStates[m.currentState]
	if !ok {
		return nil, &InvalidStateError[Tstate]{State: m.currentState}
	}

	if !m.settings.AllowSameStateTransition && m.currentState == destinationState {
		return nil, &SameStateTransitionError[Tevent]{Event: event}
	}

	return &Transition[Tstate, Tevent]{
		Source:      m.currentState,
		Destination: destinationState,
		Event:       event,
	}, nil
}

// Error returns the error message.
func (e *InvalidEventError[Tevent]) Error() string {
	return fmt.Sprintf("invalid event: %v", e.Event)
}

// Error returns the error message.
func (e *InvalidStateError[Tstate]) Error() string {
	return fmt.Sprintf("invalid state: %v", e.State)
}

// Error returns the error message.
func (e *SameStateTransitionError[Tevent]) Error() string {
	return fmt.Sprintf("same state transition not allowed: %v", e.Event)
}
