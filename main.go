package main

import (
	"fmt"

	"github.com/beam/core/pkg/states"
)

func main() {
	var st states.WorkflowState[*states.Task]
	st = states.NewTask("", "", "", false)
	fmt.Println(st.End())
}
