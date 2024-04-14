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
