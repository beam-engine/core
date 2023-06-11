package runtime

type context interface {
	Id() string

	ProcessId() string

	Name() string

	Type() string
}

type Log interface {
	WorkflowId() int64

	TaskName() string

	URL() string

	Type() string

	Request() string

	Response() string
}

type WorkflowContext interface {
	context

	SetVariable(name string, value any)

	GetVariable(name string) any

	HasVariable(name string) bool

	RemoveVariable(name string)

	RemoveAllVariables()

	GetAllVariables() map[string]any
}

type TaskContext interface {
	context

	IsLogsExist() bool

	AppendErrors(errors ...string)

	AppendLog(log Log)

	ClearLogs()

	GetLogs() []Log
}
