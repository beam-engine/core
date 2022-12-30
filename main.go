package main

import (
	"fmt"

	"github.com/beam/core/pkg/states"
)

func main() {
	dict := make(map[string]states.WorkflowState)

	dict["1"] = states.NewTask("Fun1", "", "", false)
	dict["2"] = states.NewWait("Fun2", "", "", false)
	dict["3"] = states.NewChoice(nil, "Fun3", "", false)

	for _, value := range dict {
		tp := value.Type()
		switch tp {
		case states.TypeTask:
			tk := value.(*states.Task)
			fmt.Println("Task = ", tk.Name())
		case states.TypeWait:
			wt := value.(*states.Wait)
			fmt.Println("Task = ", wt.Name())
		case states.TypeCond:
			cd := value.(*states.Choice)
			fmt.Println("Task = ", cd.Name(), cd.Conditions())
		default:
			fmt.Println("Invalid state")
		}
	}

}
