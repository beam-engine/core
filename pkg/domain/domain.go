package domain

import "github.com/beam/core/pkg/states"

type WorkflowGraph struct {
	WorkflowName   string
	IsAsync        bool
	StartAt        string
	ResultVariable string
	StatesMap      map[string]states.WorkflowState
}
