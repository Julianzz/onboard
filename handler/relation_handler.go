package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// Get relation get handler
func (handler *RelationsHandler) Get(w http.ResponseWriter, r *http.Request, params map[string]string) {

	// be cautious
	userID := string(params["user_id"])
	relations, err := model.GetRelationsByUserID(userID)
	if err != nil {
		// process
	}

	results := make([]*RelationResult, 0)
	for _, rel := range relations {
		result := &RelationResult{
			UserID: rel.WipeUserID,
			State:  rel.State,
			Type:   rel.Type,
		}
		results = append(results, result)
	}

	values, _ := json.Marshal(results)
	w.Write(values)
}

// Put relation put handler
func (handler *RelationsHandler) Put(w http.ResponseWriter, r *http.Request, params map[string]string, body []byte) {

	if err := relationSchema.Validate(bytes.NewReader(body)); err != nil {
		v := fmt.Sprintf("error %v", err)
		w.Write([]byte(v))
		return
	}
	var relation RelationResult
	if err := json.Unmarshal(body, &relation); err != nil {
		v := fmt.Sprintf("error in parse body: %v", err)
		w.Write([]byte(v))
		return
	}

	userID := string(params["user_id"])
	wipeUserID := string(params["wipe_user_id"])
	if userID == wipeUserID {
		return
	}

	user, _ := model.GetUserByID(userID)
	wipeUser, _ := model.GetUserByID(wipeUserID)
	if user == nil || wipeUser == nil {
		w.Write([]byte("user not exists"))
		return
	}

	err := model.CreateUserRelation(userID, wipeUserID, "relationship", relation.State)
	if err != nil {
	}

	rel, err := model.GetRelationsByUserIDs(userID, wipeUserID)
	relation = RelationResult{
		UserID: rel.WipeUserID,
		State:  rel.State,
		Type:   rel.Type,
	}

	values, _ := json.Marshal(relation)
	w.Write(values)
}
