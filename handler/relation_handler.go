package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/p1cn/onboard/liuzhenzhong/model"
	"github.com/santhosh-tekuri/jsonschema"
)

const userRelationSchemaJSON = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"properties": {
		"state": {
			"description": "state",
			"type": "string",
			"pattern": "^disliked|liked$"
		}
	},
	"required": ["state"]
    
}
`

var relationSchema *jsonschema.Schema

func init() {
	relationURL := "user_relation.json"
	err := compiler.AddResource(relationURL, strings.NewReader(userRelationSchemaJSON))
	if err != nil {
		log.Printf("error in load user schema %v", err)
		panic("wrong load schema")
	}

	relationSchema, err = compiler.Compile(relationURL)
	if err != nil {
		log.Printf("error in compile user schema %v", err)
		panic("error in compile user schema")
	}
}

// RelationsHandler process relation
type RelationsHandler struct {
	DefaultRestHandler
}

// RelationResult user for relation result output
type RelationResult struct {
	UserID string `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}

// build the result return
func (handler *RelationsHandler) buildRelation(rel *model.Relation) *RelationResult {
	state := rel.State
	if rel.MatchState == model.MatchedState {
		state = model.MatchedState
	}
	result := &RelationResult{
		UserID: rel.WipeUserID,
		State:  state,
		Type:   rel.Type,
	}
	return result
}

// Get relation get handler
func (handler *RelationsHandler) Get(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {

	// be cautious
	userID := string(params["user_id"])

	// check user
	user, _ := model.GetUserByID(userID)
	if user == nil {
		return nil, NewRestfulError(errors.New("can not find user"), http.StatusNotFound, "can not find user")
	}

	relations, err := model.GetRelationsByUserID(userID)
	if err != nil {
		return nil, NewRestfulError(err, http.StatusInternalServerError, "find relation wrong")
	}

	results := make([]*RelationResult, 0)
	for _, rel := range relations {
		result := handler.buildRelation(&rel)
		results = append(results, result)
	}

	return results, nil
}

// Put relation put handler
func (handler *RelationsHandler) Put(w http.ResponseWriter, r *http.Request, params map[string]string, body []byte) (interface{}, error) {

	if err := relationSchema.Validate(bytes.NewReader(body)); err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "wrong in request params")
	}
	relation := &RelationResult{}
	if err := json.Unmarshal(body, &relation); err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in parsing body")
	}

	userID := string(params["user_id"])
	wipeUserID := string(params["wipe_user_id"])
	if userID == wipeUserID {
		return nil, NewRestfulError(errors.New("same user id can not be matched"), http.StatusBadRequest, "wrong in inputing user_id")
	}

	user, _ := model.GetUserByID(userID)
	wipeUser, _ := model.GetUserByID(wipeUserID)
	if user == nil || wipeUser == nil {
		return nil, NewRestfulError(errors.New("can not find user"), http.StatusNotFound, "can not for user to update")
	}

	err := model.CreateUserRelation(userID, wipeUserID, "relationship", relation.State)
	if err != nil {
		return nil, NewRestfulError(err, http.StatusInternalServerError, "error in creating relation")
	}

	rel, err := model.GetRelationsByUserIDs(userID, wipeUserID)
	relation = handler.buildRelation(rel)

	return relation, nil
}
