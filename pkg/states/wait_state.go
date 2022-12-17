package states

type Wait struct {
	name          string
	previousState string
	nextState     string
	end           bool
}

func NewWait(name, previous, next string, end bool) *Wait {
	return &Wait{name, previous, next, end}
}

func (ws *Wait) Name() string {
	return ws.name
}

func (ws *Wait) Type() string {
	return TypeWait
}

func (ws *Wait) Previous() string {
	return ws.previousState
}

func (ws *Wait) Next() string {
	return ws.nextState
}

func (ws *Wait) End() bool {
	return ws.end
}

func (ws *Wait) Get() *Wait {
	return ws
}
