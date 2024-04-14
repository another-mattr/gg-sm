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
