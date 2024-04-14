package traffic_light

import (
	"testing"

	"github.com/another-mattr/gg-sm/ggsm"
)

func TestTrafficLightTransitions(t *testing.T) {
	tests := []struct {
		name     string
		initial  ggsm.State
		event    ggsm.Event
		expected ggsm.State
	}{
		{"Red to Green", Red, ggsm.Event{Type: Change}, Green},
		{"Green to Yellow", Green, ggsm.Event{Type: Change}, Yellow},
		{"Yellow to Red", Yellow, ggsm.Event{Type: Change}, Red},
		{"Red to BlinkRed", Red, ggsm.Event{Type: PauseService, Payload: PauseServicePayload{Reason: "routine maintenance"}}, BlinkRed},
		{"BlinkRed to BlinkRed", BlinkRed, ggsm.Event{Type: Change}, BlinkRed},
		{"BlinkRed to Red", BlinkRed, ggsm.Event{Type: ResumeService}, Red},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := NewTrafficLight(tt.name, tt.initial)

			tl.HandleEvent(tt.event)

			if tl.GetCurrentState() != tt.expected {
				t.Errorf("got %v, want %v", tl.GetCurrentState(), tt.expected)
			}
		})
	}
}
