package parser

import (
	"errors"
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

// Error Messages
const ERR_INVALID_JSON_FILE = "Cannot able to create workflow graph from Json"
const ERR_REQUIRED_INFO_MISSING = "required information in workflow Json is missing"
const ERR_INVALID_WORKFLOW = "The provided workflow json is wrong"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

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

	ue := gjson.Get(json, "name.last")
	println(ue.String())
	return nil, nil
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

	log.Info().Msg("return from validateFields")
	return nil
}
