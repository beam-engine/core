package states

type WaitState struct {
	name     string
	previous string
	next     string
	end      bool
}

func NewWait(name, previous, next string, end bool) *WaitState {
	return &WaitState{name, previous, next, end}
}

func (ws *WaitState) Name() string {
	return ws.name
}

func (ws *WaitState) Type() StateType {
	return Wait
}

func (ws *WaitState) Previous() string {
	return ws.previous
}

func (ws *WaitState) Next() string {
	return ws.next
}

func (ws *WaitState) End() bool {
	return ws.end
}
