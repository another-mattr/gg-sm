package traffic_light

import (
	"fmt"

	"github.com/another-mattr/gg-sm/ggsm"
)

const (
	Change        ggsm.EventType = "Change"
	PauseService  ggsm.EventType = "PauseService"
	ResumeService ggsm.EventType = "ResumeService"
)

type ChangeEventPayload struct {
	Count int
}

type PauseServicePayload struct {
	Reason string
}

const (
	Red      ggsm.State = "Red"
	Yellow   ggsm.State = "Yellow"
	Green    ggsm.State = "Green"
	BlinkRed ggsm.State = "BlinkRed"
)

type TrafficLight struct {
	id string
	ggsm.StateMachine
}

func NewTrafficLight(id string, currentState ggsm.State) *TrafficLight {
	tl := &TrafficLight{ id: id }
	tl.Initialize(currentState)
	tl.AddTransition(Red, Change, Green)
	tl.AddTransition(Red, PauseService, BlinkRed)
	tl.AddTransition(Yellow, Change, Red)
	tl.AddTransition(Green, Change, Yellow)
	tl.AddTransition(BlinkRed, Change, BlinkRed)
	tl.AddTransition(BlinkRed, ResumeService, Red)

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

	tl.AddOnExit(Green, func(e ggsm.Event) error {
		fmt.Printf("Leaving Green: %v\n", e.Payload)
		return nil
	})

	tl.AddOnEnter(BlinkRed, func(e ggsm.Event) error {
		switch e.Payload.(type) {
		case ChangeEventPayload:
			p := e.Payload.(ChangeEventPayload)
			fmt.Printf("Blink Red. %v\n", p.Count)
			if p.Count > 10 {
				tl.HandleEvent(ggsm.Event{Type: ResumeService, Payload: nil})
			}
		case PauseServicePayload:
			p := e.Payload.(PauseServicePayload)
			fmt.Printf("Blink Red. Service paused because: %s\n", p.Reason)
		default:
			return fmt.Errorf("invalid payload type: %T", e.Payload)
		}
		return nil
	})

	return tl
}

// GetId: Returns the ID of the traffic light.
func (tl *TrafficLight) GetId() string {
	return tl.id
}