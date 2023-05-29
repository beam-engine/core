package main

import (
	"fmt"

	"github.com/beam/core/pkg/states"
)

func main() {
	dict := make(map[string]states.WorkflowState)

	dict["1"] = states.NewTask("Fun1", "", "", false)
	dict["2"] = states.NewWait("Fun2", "", "", false)
	dict["3"] = states.NewChoice(nil, "Fun3", "", "", false)

	for _, value := range dict {
		tp := value.Type()
		switch tp {
		case states.Task:
			tk := value.(*states.TaskState)
			fmt.Println("Task Name = ", tk.Name())
		case states.Wait:
			wt := value.(*states.WaitState)
			fmt.Println("Wait Name = ", wt.Name())
		case states.Cond:
			cd := value.(*states.ChoiceState)
			fmt.Println("Condition Name = ", cd.Name(), cd.Conditions())
		default:
			fmt.Println("Invalid state")
		}
	}

}
