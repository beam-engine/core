package states

type TaskState struct {
	name     string
	previous string
	next     string
	end      bool
}

func NewTask(name, previous, next string, end bool) *TaskState {
	return &TaskState{name, previous, next, end}
}

func (ts *TaskState) Name() string {
	return ts.name
}

func (ts *TaskState) Type() StateType {
	return Task
}

func (ts *TaskState) Previous() string {
	return ts.previous
}

func (ts *TaskState) Next() string {
	return ts.next
}

func (ts *TaskState) End() bool {
	return ts.end
}
