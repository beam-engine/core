package runtime

import "errors"

type ApplicationContext struct {
	taskMap map[string]*WorkflowTask
}

func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{
		taskMap: map[string]*WorkflowTask{},
	}
}

func (app *ApplicationContext) RegisterTask(name string, task *WorkflowTask) {
	app.taskMap[name] = task
}

func (app *ApplicationContext) GetTask(name string) (*WorkflowTask, error) {
	if instance, ok := app.taskMap[name]; ok {
		return instance, nil
	} else {
		return nil, errors.New("task " + name + " not found in the application context")
	}
}
