package domain

import (
	"errors"
	"github.com/beam/core/pkg/states"
)

type OrderedMap struct {
	orderedKeys []string
	states      map[string]*states.WorkflowState
	startIndex  int
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		orderedKeys: make([]string, 0),
		states:      map[string]*states.WorkflowState{},
		startIndex:  0,
	}
}

func (om *OrderedMap) Insert(component string, state *states.WorkflowState) error {
	if _, ok := om.states[component]; !ok {
		return errors.New("component already exists, cannot update the existing component")
	}
	om.orderedKeys[om.startIndex] = component
	om.startIndex += 1

	om.states[component] = state
	return nil
}
