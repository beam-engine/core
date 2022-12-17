package states

type Task struct {
	name          string
	previousState string
	nextState     string
	end           bool
}

func NewTask(name, previous, next string, end bool) *Task {
	return &Task{name, previous, next, end}
}

func (ts *Task) Name() string {
	return ts.name
}

func (ts *Task) Type() string {
	return TypeTask
}

func (ts *Task) Previous() string {
	return ts.previousState
}

func (ts *Task) Next() string {
	return ts.nextState
}

func (ts *Task) End() bool {
	return ts.end
}

func (ts *Task) Get() *Task {
	return ts
}
