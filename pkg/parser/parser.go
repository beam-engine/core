package parser

import (
	"errors"
	"github.com/beam/core/pkg/states"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"

	"github.com/beam/core/pkg/domain"
)

// LabelName Name of the workflow
const LabelName = "Name"

// LabelDescription Description in workflow JSON
const LabelDescription = "Description"

// LabelInitState The First State in the workflow
const LabelInitState = "InitState"

// LabelAsync Is workflow needs to executed in async mode
const LabelAsync = "Async"

// LabelStates All workflow states
const LabelStates = "States"

// LabelMode workflow engine mode Realtime, Near-Realtime and Non-Realtime.
const LabelMode = "Mode"

// LabelResultVariable Result Variable
const LabelResultVariable = "ResultVariable"

// LabelType Workflow Type
const LabelType = "Type"

const TypeTask = "Task"
const TypeCondition = "Condition"

const TypeWait = "Wait"

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

	workflowName := jsonNode.Get(LabelName).Str

	startState := jsonNode.Get(LabelInitState).Str

	isAsync := jsonNode.Get(LabelAsync).Bool()

	resultVariable := ""
	if !isAsync {
		resultVariable = jsonNode.Get(LabelResultVariable).Str
	}

	mode := jsonNode.Get(LabelMode).Str
	engineMode, err := transFormMode(mode)

	if err != nil {
		log.Error().Msg("error: problem while parsing mode value in workflow json - " + err.Error())
		return nil, err
	}

	stateNode := jsonNode.Get(LabelStates)

	stateNode.ForEach(func(key, value gjson.Result) bool {
		component := key.String()
		stateInstance, err := jp.createWorkflowState(value)
		if err != nil {
			log.Error().Msg("error: cannot able to create workflow state for " + component)
			os.Exit(1)
		}
		workflowStateMap[component] = stateInstance
		return true
	})

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
	stateType := stateNode.Get(LabelType).String()

	var result states.WorkflowState
	var err error = nil

	switch stateType {
	case TypeTask:
		result, err = jp.createTaskState(stateNode)
		break
	case TypeWait:
		result, err = jp.createWaitState(stateNode)
		break
	case TypeCondition:
		result, err = jp.createChoiceState(stateNode)
		break
	default:
		err = errors.New("unknown state type found " + stateType)
		break
	}

	return &result, err
}

func (jp *jsonParser) createTaskState(stateNode gjson.Result) (*states.TaskState, error) {
	return &states.TaskState{}, nil
}

func (jp *jsonParser) createWaitState(stateNode gjson.Result) (*states.WaitState, error) {
	return &states.WaitState{}, nil
}

func (jp *jsonParser) createChoiceState(stateNode gjson.Result) (*states.ChoiceState, error) {
	return &states.ChoiceState{}, nil
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

	if !jsonNode.Get(LabelName).Exists() {
		log.Error().Msg("error: required label" + LabelName + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelName)
	}

	if !jsonNode.Get(LabelDescription).Exists() {
		log.Error().Msg("error: required label" + LabelDescription + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelDescription)
	}

	if !jsonNode.Get(LabelAsync).Exists() {
		log.Error().Msg("error: required label" + LabelAsync + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelAsync)
	}

	if !jsonNode.Get(LabelStates).Exists() {
		log.Error().Msg("error: required label" + LabelStates + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelStates)
	}

	if !jsonNode.Get(LabelInitState).Exists() {
		log.Error().Msg("error: required label" + LabelInitState + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelInitState)
	}

	if !jsonNode.Get(LabelMode).Exists() {
		log.Error().Msg("error: required label" + LabelInitState + " not found in json")
		return errors.New(ERR_REQUIRED_INFO_MISSING + "= " + LabelInitState)
	}

	log.Info().Msg("return from validateFields")
	return nil
}
