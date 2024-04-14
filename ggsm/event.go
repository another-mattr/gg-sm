package ggsm

import (
	"fmt"
	"github.com/another-mattr/gg-sm/types"
)

// EventType: The type of event that triggers a state transition in the state machine.
type EventType string

// Event: A struct that represents an event that triggers a state transition in the state machine.
type Event struct {
	Type EventType
	Payload any
}

// GetType: Returns the type of the event.
func (e Event) GetType() EventType {
	return e.Type
}

// ExtractPayload: Extracts the payload from the event and returns it as a pointer to the specified type.
// If the payload is not of the specified type, an error is returned.
func ExtractPayload[T any](e Event) (*T, error) {
	v, ok := e.Payload.(T)
	if !ok {
		expectedTypeZeroValue := types.GetZeroValue[T]()
		return nil, fmt.Errorf("invalid payload type. Expected: %T, Received: %T", expectedTypeZeroValue, e.Payload)
	}
	return &v, nil
}