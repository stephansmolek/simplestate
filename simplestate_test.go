package simplestate

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, testStateA, m.currentState)
		assert.Equal(t, 3, len(m.events))
		assert.Equal(t, settings, m.settings)
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
	if m != nil {
		assert.Equal(t, testStateA, m.Get())
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
		assert.True(t, m.Is(testStateA))
		assert.False(t, m.Is(testStateB))
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
		assert.Equal(t, testStateB, m.Get())
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
		assert.True(t, m.Can(testEventId1))
		assert.False(t, m.Can(testEventId2))
		assert.False(t, m.Can(testEventId3))
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
		assert.False(t, m.Can("invalid-event"))
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
		assert.False(t, m.Can(testEventId2))
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
		assert.True(t, m.Can(testEventId3))
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
		assert.True(t, m.Can(testEventId1))
		assert.NoError(t, m.Trigger(testEventId1))
		assert.Equal(t, testStateC, m.Get())
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
		assert.False(t, m.Can("invalid-event"))
		assert.Error(t, m.Trigger("invalid-event"))
		assert.Equal(t, testStateA, m.Get())
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
		assert.False(t, m.Can(testEventId2))
		assert.Error(t, m.Trigger(testEventId2))
		assert.Equal(t, testStateA, m.Get())
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
		assert.True(t, m.Can(testEventId3))
		assert.NoError(t, m.Trigger(testEventId3))
		assert.Equal(t, testStateA, m.Get())
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
		assert.False(t, m.Can(testEventId3))
		assert.Error(t, m.Trigger(testEventId3))
		assert.Equal(t, testStateA, m.Get())
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
		assert.False(t, m.Can(testEventId2))
		assert.Error(t, m.Trigger(testEventId2))
		assert.Equal(t, testStateA, m.Get())
	}
}

func TestInvaildEventError(t *testing.T) {
	err := &InvalidEventError[TestEvent]{"invalid-event"}
	assert.Equal(t, "invalid event: invalid-event", err.Error())
}

func TestInvalidStateError(t *testing.T) {
	err := &InvalidStateError[TestState]{testStateA}
	assert.Equal(t, "invalid state: state-A", err.Error())
}

func TestSameStateTransitionError(t *testing.T) {
	err := &SameStateTransitionError[TestState]{testStateA}
	assert.Equal(t, "same state transition not allowed: state-A", err.Error())
}

func TestTriggerFromCallback(t *testing.T) {
	settings := Settings{
		AllowSameStateTransition: false,
	}
	var m *Machine[TestState, TestEvent]
	callback := func(transition Transition[TestState, TestEvent]) {
		if transition.Event == testEventId1 {
			assert.NoError(t, m.Trigger(testEventId2))
		}
	}
	m = NewMachine(testStateA, testEvents, settings, callback)
	if m == nil {
		t.Error("NewMachine returned nil")
	}
	if m != nil {
		assert.True(t, m.Can(testEventId1))
		assert.NoError(t, m.Trigger(testEventId1))
		assert.Equal(t, testStateD, m.Get())
	}
}
