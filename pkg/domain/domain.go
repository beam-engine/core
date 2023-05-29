package domain

import "github.com/beam/core/pkg/states"

type Engine int

const (
	ExpressEngine Engine = iota
	DefaultEngine
)

type WorkflowGraph struct {
	States         map[string]states.WorkflowState
	WorkflowName   string
	StartAt        string
	ResultVariable string
	Mode           Engine
	IsAsync        bool
}
