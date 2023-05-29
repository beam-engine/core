package states

type Expression struct {
	Variable string
	Value    string
	Type     MatchType
}

func NewExpression(variable string, matchType MatchType, matchValue string) *Expression {
	return &Expression{Variable: variable, Type: matchType, Value: matchValue}
}

// Condition interface contract
type Condition struct {
	Expressions []Expression
	Next        string
	Type        ConditionType
}

func NewCondition(expressions []Expression, conditionType ConditionType, next string) *Condition {
	return &Condition{Expressions: expressions, Type: conditionType, Next: next}
}

type ChoiceState struct {
	conditions []*Condition
	name       string
	previous   string
	next       string
	end        bool
}

func NewChoice(conditionList []*Condition, name, previous, next string, end bool) *ChoiceState {
	return &ChoiceState{conditionList, name, previous, next, end}
}

func (cs *ChoiceState) Name() string {
	return cs.name
}

func (cs *ChoiceState) Type() StateType {
	return Cond
}

func (cs *ChoiceState) Previous() string {
	return cs.previous
}

func (cs *ChoiceState) Next() string {
	return cs.next
}

func (cs *ChoiceState) End() bool {
	return cs.end
}

func (cs *ChoiceState) Conditions() []*Condition {
	return cs.conditions
}
