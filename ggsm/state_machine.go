// WARNING: This file has been generated. Do not attempt to modify it manually.

package ggsm

import (
	"fmt"
)

// State: A node's name in our state machine.
type State string

// TransitionMap: A map of maps that represents the transitions between states based on events.
type TransitionMap map[State]map[EventType]State

// Callback: The callback type that is executed when a state transition occurs.
type Callback func(e Event) error

// CallbackMap: A map of callbacks that are executed when the state machine enters or exits a specific state.
type CallbackMap map[State]Callback

type StateMachine struct {
	currentState State
	transitions  TransitionMap
	onEnter      CallbackMap
	onExit       CallbackMap
}

// Initialize sets the initial state of the state machine.
func (sm *StateMachine) Initialize(state State) error {
	if sm.currentState != "" {
		return fmt.Errorf("initial state already set")
	}
	sm.currentState = state
	return nil
}

// GetCurrentState returns the current state of the state machine.
func (sm *StateMachine) GetCurrentState() State {
	return sm.currentState
}

// AddTransition configures a transition from one state to another based on a given event.
// You can think of it like a directed graph: (from)-[Event]->(to)
func (sm *StateMachine) AddTransition(from State, event EventType, to State) {
	if sm.transitions == nil {
		sm.transitions = make(TransitionMap)
	}
	if _, ok := sm.transitions[from]; !ok {
		sm.transitions[from] = make(map[EventType]State)
	}
	sm.transitions[from][event] = to
}

// AddOnEnter adds a callback to be executed when the state machine enters a specific state.
func (sm *StateMachine) AddOnEnter(state State, callback Callback) {
	if sm.onEnter == nil {
		sm.onEnter = make(CallbackMap)
	}
	sm.onEnter[state] = callback
}

// AddOnExit adds a callback to be executed when the state machine exits a specific state.
func (sm *StateMachine) AddOnExit(state State, callback Callback) {
	if sm.onExit == nil {
		sm.onExit = make(CallbackMap)
	}
	sm.onExit[state] = callback
}

// HandleEvent handles an event by transitioning the state machine to the next state based on the event.
func (sm *StateMachine) HandleEvent(event Event) {
	if transitions, ok := sm.transitions[sm.currentState]; ok {
		if nextState, ok := transitions[event.GetType()]; ok {
			// Execute the onExit callback for the current state, if defined.
			if callback, ok := sm.onExit[sm.currentState]; ok {
				callback(event)
			}

			// Update the current state to the next state.
			sm.currentState = nextState

			// Execute the onEnter callback for the next state, if defined.
			if callback, ok := sm.onEnter[nextState]; ok {
				callback(event)
			}
		} else {
			fmt.Printf("No transitions found for the event %s in the current state %s\n", event.GetType(), sm.currentState)
		}
	} else {
		fmt.Printf("No transitions found for the current state %s\n", sm.currentState)
	}
}
