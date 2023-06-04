package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/beam/core/pkg/parser"
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

	file, err := os.Open("/Users/kishorekarunakaran/coding2fun/beam-engine/core/resources/RE2RE.json")
	if err != nil {
		log.Error().Msg("Cannot Read File = " + err.Error())
		os.Exit(0)
	}
	defer file.Close()

	p := parser.NewJsonParser()
	res, _ := p.CreateWorkflowGraph(file)
	fmt.Println(res)
}
