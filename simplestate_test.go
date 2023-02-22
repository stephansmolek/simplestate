package simplestate

import (
	"testing"
)

var (
	testStateA TestState = "state-A"
	testStateB TestState = "state-B"
	testStateC TestState = "state-C"
	testStateD TestState = "state-D"

	testEventId1 TestEvent = "event-1"
	testEventId2 TestEvent = "event-2"
	testEventId3 TestEvent = "event-3"

	testEvent1 = Event[TestState, TestEvent]{
		Event: testEventId1,
		Sources: []TestState{
			testStateA,
			testStateB,
		},
		Destination: testStateC,
	}

	testEvent2 = Event[TestState, TestEvent]{
		Event: testEventId2,
		Sources: []TestState{
			testStateC,
		},
		Destination: testStateD,
	}

	testEvent3 = Event[TestState, TestEvent]{
		Event: testEventId3,
		Sources: []TestState{
			testStateA,
			testStateD,
		},
		Destination: testStateA,
	}

	testEvents = Events[TestState, TestEvent]{
		testEvent1,
		testEvent2,
		testEvent3,
	}
)

type (
	TestState string
	TestEvent string
)

func TestNewMachine(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.currentState != testStateA {
			t.Errorf("NewMachine returned invalid state: %v", m.currentState)
		}

		if len(m.events) != 3 {
			t.Errorf("NewMachine returned invalid events: %v", m.events)
		}

		if m.settings != settings {
			t.Errorf("NewMachine returned invalid settings: %v", m.settings)
		}
	}
}

func TestNewMachineWithInvalidEvent(t *testing.T) {
}

func TestGet(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil && m.Get() != testStateA {
		t.Errorf("Get() returned invalid state: %v", m.Get())
	}
}

func TestIs(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Is(testStateA) {
			t.Errorf("Is() returned invalid state: %v", m.Is(testStateA))
		}
		if m.Is(testStateB) {
			t.Errorf("Is() returned invalid state: %v", m.Is(testStateB))
		}
	}
}

func TestSet(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		m.Set(testStateB)
		if m.Get() != testStateB {
			t.Errorf("Invalid state after Set(): %v", m.Get())
		}
	}
}

func TestCan(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Can(testEventId1) {
			t.Errorf("Can() true expected for event: %v", testEventId1)
		}

		if m.Can(testEventId2) {
			t.Errorf("Can() false expected for event: %v", testEventId2)
		}

		if m.Can(testEventId3) {
			t.Errorf("Can() false expected for event: %v", testEventId3)
		}
	}
}

func TestCanWithInvalidEvent(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can("invalid-event") {
			t.Errorf("Can() false expected for invalid event")
		}
	}
}

func TestCanWithInvalidState(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can(testEventId2) {
			t.Errorf("Can() false expected for event: %v", testEventId2)
		}
	}
}

func TestCanWithSameStateTransition(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: true,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Can(testEventId3) {
			t.Errorf("Can() true expected for event: %v", testEventId3)
		}
	}
}

func TestTrigger(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Can(testEventId1) {
			t.Errorf("Can() true expected for event: %v", testEventId1)
		}

		if m.Trigger(testEventId1) != nil {
			t.Errorf("Trigger() returned error for event: %v", testEventId1)
		}

		if m.Get() != testStateC {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestTriggerWithInvalidEvent(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can("invalid-event") {
			t.Errorf("Can() false expected for invalid event")
		}

		if m.Trigger("invalid-event") == nil {
			t.Errorf("Trigger() returned no error for invalid event")
		}

		if m.Get() != testStateA {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestTriggerWithInvalidState(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can(testEventId2) {
			t.Errorf("Can() false expected for event: %v", testEventId2)
		}

		if m.Trigger(testEventId2) == nil {
			t.Errorf("Trigger() returned no error for invalid event")
		}

		if m.Get() != testStateA {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestTriggerWithSameStateTransition(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: true,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Can(testEventId3) {
			t.Errorf("Can() true expected for event: %v", testEventId3)
		}

		if m.Trigger(testEventId3) != nil {
			t.Errorf("Trigger() returned error for event: %v", testEventId3)
		}

		if m.Get() != testStateA {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestTriggerWithSameStateTransitionDisallowed(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can(testEventId3) {
			t.Errorf("Can() false expected for event: %v", testEventId3)
		}

		if m.Trigger(testEventId3) == nil {
			t.Errorf("Trigger() returned no error for invalid event")
		}

		if m.Get() != testStateA {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestTriggerWithInvalidTransition(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	callback := func(transition Transition[TestState, TestEvent]) {}
	m := NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if m.Can(testEventId2) {
			t.Errorf("Can() false expected for event: %v", testEventId2)
		}

		if m.Trigger(testEventId2) == nil {
			t.Errorf("Trigger() returned no error for invalid event")
		}

		if m.Get() != testStateA {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}

func TestInvaildEventError(t *testing.T) {
	err := &InvalidEventError[TestEvent]{"invalid-event"}

	if err.Error() != "invalid event: invalid-event" {
		t.Errorf("InvalidEventError returned invalid error: %v", err.Error())
	}
}

func TestInvalidStateError(t *testing.T) {
	err := &InvalidStateError[TestState]{testStateA}

	if err.Error() != "invalid state: state-A" {
		t.Errorf("InvalidStateError returned invalid error: %v", err.Error())
	}
}

func TestSameStateTransitionError(t *testing.T) {
	err := &SameStateTransitionError[TestState]{testStateA}

	if err.Error() != "same state transition not allowed: state-A" {
		t.Errorf("SameStateTransitionError returned invalid error: %v", err.Error())
	}
}

func TestTriggerFromCallback(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	var m *Machine[TestState, TestEvent]
	callback := func(transition Transition[TestState, TestEvent]) {
		if transition.Event == testEventId1 {
			if m.Trigger(testEventId2) != nil {
				t.Errorf("Trigger() returned error for event: %v", testEventId2)
			}
		}
	}
	m = NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		if !m.Can(testEventId1) {
			t.Errorf("Can() true expected for event: %v", testEventId1)
		}

		if m.Trigger(testEventId1) != nil {
			t.Errorf("Trigger() returned error for event: %v", testEventId1)
		}

		if m.Get() != testStateD {
			t.Errorf("Invalid state after Trigger(): %v", m.Get())
		}
	}
}
