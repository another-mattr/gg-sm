# gg-sm
Go Go State Machine

A super-simple state machine library for Go. It supports basic transitions as well as callbacks for entering and exiting states.

## Usage

Here's a simple traffic light example:

```go
// filename = main.go
package main

import (
	"github.com/another-mattr/gg-sm/examples/simple_traffic_light/traffic_light"
	"github.com/another-mattr/gg-sm/ggsm"
)

func main() {
	tl := traffic_light.NewTrafficLight("main & oak", traffic_light.Red)
	for i := 0; i < 6; i++ {
		tl.HandleEvent(ggsm.Event{
			Type:    traffic_light.Change,
			Payload: traffic_light.ChangeEventPayload{Count: i},
		})
	}
}
```

```go
// filename = traffic_light.go
package traffic_light

import (
	"fmt"

	"github.com/another-mattr/gg-sm/ggsm"
)

const (
	Change ggsm.EventType = "Change"
)

type ChangeEventPayload struct {
	Count int
}

const (
	Red    ggsm.State = "Red"
	Yellow ggsm.State = "Yellow"
	Green  ggsm.State = "Green"
)

type TrafficLight struct {
	id string
	ggsm.StateMachine
}

func NewTrafficLight(id string, currentState ggsm.State) *TrafficLight {
	tl := &TrafficLight{id: id}
	tl.Initialize(currentState)
	tl.AddTransition(Green, Change, Yellow)
	tl.AddTransition(Yellow, Change, Red)
	tl.AddTransition(Red, Change, Green)

	tl.AddOnEnter(Red, func(e ggsm.Event) error {
		fmt.Printf("Red: %v\n", e.Payload)
		return nil
	})

	tl.AddOnEnter(Yellow, func(e ggsm.Event) error {
		fmt.Printf("Yellow: %v\n", e.Payload)
		return nil
	})

	tl.AddOnEnter(Green, func(e ggsm.Event) error {
		fmt.Printf("Green: %v\n", e.Payload)
		return nil
	})

	return tl
}

// GetId: Returns the ID of the traffic light.
func (tl *TrafficLight) GetId() string {
	return tl.id
}
```