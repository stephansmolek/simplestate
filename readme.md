# Simple state machine
## Overview
This is a simple state machine with generic types for event and state identifiers. Possible state transitions as well as a start state must be given. Same state transitions can be allowed or forbidden depending upon configuration. A callback after a transition is available.

## How to use
1. Create state and event types, for example:
```go
type (
	ExampleState string
	ExampleEvent string
)
```

2. Define available state identifiers:
```go
var (
    exampleStateA1 ExampleState = "example-state-A1"
    exampleStateA2 ExampleState = "example-state-A2"
    exampleStateB ExampleState = "example-state-B"
    exampleStateC ExampleState = "example-state-C"
)
```

3. Define available event identifiers:
```go
var (
    exampleEventA ExampleEvent = "example-event-A"
    exampleEventB ExampleEvent = "example-event-B"
)
```

4. Define available events:
```go
var events = []Event[ExampleState, ExampleEvent] {
    {
        Event: exampleEventA,
        Sources: []ExampleState {
            exampleStateA1,
            exampleStateA2,
        },
        Destination: exampleStateB,
    },
    {
        Event: exampleEventB,
        Sources: []ExampleState {
            exampleStateB,
        },
        Destination: exampleStateC,
    },
}
```

5. Create settings:
```go
var settings = Settings{
		AllowSameStateTransition: false,
	}
```

6. Define callback function:
```go
var callback = func(transition Transition[ExampleState, ExampleEvent]) {
		...
	}
```

7. Create state machine:
```go
var machine = NewMachine(exampleStateA1, events, settings, callback)
```

8. Start triggering events:
```go
err := machine.Trigger(exampleEventA); if err != nil {
    ...
}
```
