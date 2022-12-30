package states

// State Type
const (
	TypeTask = "Task"
	TypeCond = "Condition"
	TypeWait = "Wait"
	TypePar  = "Parallel"
)

// Condition Type
const (
	Simple = "Simple"
	And    = "And"
	Or     = "Or"
)

// Match Type
const (
	StringEquals    = "StringEquals"
	StringNotEquals = "StringNotEquals"
)

type WorkflowState interface {
	Name() string
	Type() string
	Previous() string
	Next() string
	End() bool
}
