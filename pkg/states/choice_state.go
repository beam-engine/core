package states

type Expression struct {
	Variable   string
	MatchType  string
	MatchValue string
}

func NewExpression(variable, matchType, matchValue string) *Expression {
	return &Expression{Variable: variable, MatchType: matchType, MatchValue: matchValue}
}

// Condition interface contract
type Condition struct {
	Expressions   []Expression
	ConditionType string
	Next          string
}

func NewCondition(expressions []Expression, conditionType, next string) *Condition {
	return &Condition{Expressions: expressions, ConditionType: conditionType, Next: next}
}

type Choice struct {
	conditions []*Condition
	previous   string
	next       string
	end        bool
}

func NewChoice(conditionList []*Condition, previous, next string, end bool) *Choice {
	return &Choice{conditionList, previous, next, end}
}

func (cs *Choice) Name() string {
	return "Condition_No_Name"
}

func (cs *Choice) Type() string {
	return TypeCond
}

func (cs *Choice) Previous() string {
	return cs.previous
}

func (cs *Choice) Next() string {
	return cs.next
}

func (cs *Choice) End() bool {
	return cs.end
}

func (cs *Choice) Conditions() []*Condition {
	return cs.conditions
}
