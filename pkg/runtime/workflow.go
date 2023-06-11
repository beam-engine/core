package runtime

type WorkflowTask interface {
	Execute(w *WorkflowContext, t TaskContext)

	OnStateTransition(w *WorkflowContext, t TaskContext) StateTransitionResult
}

type StateTransitionResult struct {
	IsSuccess bool
}
