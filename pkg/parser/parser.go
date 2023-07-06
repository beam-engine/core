package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beam/core/pkg/states"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"

	"github.com/beam/core/pkg/domain"
)

// Name of the workflow
const LABEL_NAME = "Name"

// Description in workflow JSON
const LABEL_DESCRIPTION = "Description"

// The First State in the workflow
const LABEL_INIT_STATE = "InitState"

// Is workflow needs to executed in async mode
const LABEL_ASYNC = "Async"

// All workflow states
const LABEL_STATES = "States"

// workflow engine mode Realtime, Near-Realtime and Non-Realtime.
const LABEL_MODE = "Mode"

// Result Variable
const LABEL_RESULT_VARIABLE = "ResultVariable"

// Workflow Type
const LABEL_TYPE = "Type"

// Workflow next state
const LABEL_NEXT = "Next"

const LABEL_BRANCHES = "Branches"

// Is End of the workflow
const LABEL_END = "End"

const EXPRESSION_SIMPLE = "Simple"

const EXPRESSION_AND = "And"

const EXPRESSION_OR = "Or"

// Workflow conditions and it's related labels
const LABEL_CONDITIONS = "Conditions"
const LABEL_CONDITION_VARIABLE = "Variable"
const LABEL_CONDITION_MATCH_TYPE = "MatchType"
const LABEL_CONDITION_MATCH_VALUE = "MatchValue"
const LABEL_EXPRESSION = "Expression"
const LABEL_CONDITION_DEFAULT = "Default"

// Workflow Engine mode values
const LABEL_ENGINE_EXPRESS = "Express"
const LABEL_ENGINE_STANDARD = "Standard"

// Error Messages
const ERR_INVALID_JSON_FILE = "Cannot able to create workflow graph from Json"
const ERR_REQUIRED_INFO_MISSING = "required information in workflow Json is missing"
const ERR_INVALID_WORKFLOW = "The provided workflow json is wrong"

// Parser contract, Generates [WorkflowGraph] from workflow definition
type Parser interface {
	CreateWorkflowGraph(file *os.File) (*domain.WorkflowGraph, error)
}

type jsonParser struct {
}

func NewJsonParser() Parser {
	return &jsonParser{}
}

func (jp *jsonParser) CreateWorkflowGraph(file *os.File) (*domain.WorkflowGraph, error) {
	log.Info().Msg("in jsonParser: CreateWorkflowGraph")

	workflowDefinition, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Msg("error: unknown while parsing json - " + err.Error())
		return nil, err
	}

	jsonContent := string(workflowDefinition)
	jsonNode := gjson.Parse(jsonContent)

	if err = validateFields(jsonNode); err != nil {
		log.Error().Msg("error: validation failed, required data is missing - " + err.Error())
		return nil, err
	}

	workflowStateMap := map[string]*states.WorkflowState{}

	workflowName := jsonNode.Get(LABEL_NAME).Str

	startState := jsonNode.Get(LABEL_INIT_STATE).Str

	isAsync := jsonNode.Get(LABEL_ASYNC).Bool()

	resultVariable := ""
	if !isAsync {
		resultVariable = jsonNode.Get(LABEL_RESULT_VARIABLE).Str
	}

	mode := jsonNode.Get(LABEL_MODE).Str
	engineMode, err := transFormMode(mode)

	if err != nil {
		log.Error().Msg("error: problem while parsing mode value in workflow json - " + err.Error())
		return nil, err
	}

	stateNode := jsonNode.Get(LABEL_STATES)

	var dataMap map[string]any
	err = json.Unmarshal([]byte(stateNode.String()), &dataMap)
	if err != nil {
		log.Error().Msg("Error = " + err.Error())
		return nil, err
	}
	for key, _ := range dataMap {
		fmt.Println(key)
	}

	stateNode.ForEach(func(key, value gjson.Result) bool {
		fmt.Println("Key = ", key.String(), " and Value = ", value)
		return true
	})

	currentState := startState
	for currentState != "" && stateNode.Get(currentState).Exists() {
		stateNode := stateNode.Get(currentState)
		/*stateNode.ForEach(func(key, value gjson.Result) bool {
			fmt.Println("Key = ", key, " and Value = ", value)
			return true
		})*/
		workflowState, err := jp.createWorkflowState(stateNode)
		if err != nil {
			log.Error().Msg("unable to create workflow state from the state node " + stateNode.Str + " with the error - " + err.Error())
			return nil, err
		}
		workflowStateMap[currentState] = workflowState
		currentState = ""
	}

	workflowGraph := domain.WorkflowGraph{
		States:         nil,
		WorkflowName:   workflowName,
		StartAt:        startState,
		ResultVariable: resultVariable,
		Mode:           engineMode,
		IsAsync:        isAsync,
	}
	return &workflowGraph, nil
}

func (jp *jsonParser) createWorkflowState(stateNode gjson.Result) (*states.WorkflowState, error) {
	return nil, nil
}

func transFormMode(mode string) (domain.Engine, error) {
	log.Info().Msg("in transFormMode")
	log.Info().Msg("mode = " + mode)

	var result domain.Engine
	var err error

	switch mode {
	case LABEL_ENGINE_EXPRESS:
		result = domain.ExpressEngine
		break
	case LABEL_ENGINE_STANDARD:
		result = domain.DefaultEngine
		break
	default:
		err = errors.New("error: invalid mode type provided, expected values = [" + LABEL_ENGINE_EXPRESS + ", " + LABEL_ENGINE_STANDARD + "]")
		break
	}
	return result, err
}

func validateFields(jsonNode gjson.Result) error {
	log.Info().Msg("in validateFields")
	log.Info().Msg("workflow json = " + jsonNode.String())

	if !jsonNode.Get(LABEL_NAME).Exists() {
		log.Error().Msg("error: required label" + LABEL_NAME + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_NAME)
	}

	if !jsonNode.Get(LABEL_DESCRIPTION).Exists() {
		log.Error().Msg("error: required label" + LABEL_DESCRIPTION + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_DESCRIPTION)
	}

	if !jsonNode.Get(LABEL_ASYNC).Exists() {
		log.Error().Msg("error: required label" + LABEL_ASYNC + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_ASYNC)
	}

	if !jsonNode.Get(LABEL_STATES).Exists() {
		log.Error().Msg("error: required label" + LABEL_STATES + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_STATES)
	}

	if !jsonNode.Get(LABEL_INIT_STATE).Exists() {
		log.Error().Msg("error: required label" + LABEL_INIT_STATE + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_INIT_STATE)
	}

	if !jsonNode.Get(LABEL_MODE).Exists() {
		log.Error().Msg("error: required label" + LABEL_INIT_STATE + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LABEL_INIT_STATE)
	}

	log.Info().Msg("return from validateFields")
	return nil
}
