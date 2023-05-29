package states

type StateType int
type ConditionType int
type MatchType int

// State Type
const (
	Task StateType = iota
	Cond
	Wait
	Par
)

const (
	Simple ConditionType = iota
	And
	Or
)

// Match Type
const (
	StringEquals MatchType = iota
	StringNotEquals
)

type WorkflowState interface {
	Name() string
	Type() StateType
	Previous() string
	Next() string
	End() bool
}
