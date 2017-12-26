package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-plus/uuid"
	"github.com/santhosh-tekuri/jsonschema"

	"github.com/p1cn/onboard/liuzhenzhong/model"
)

// UsersHandler handler user request
type UsersHandler struct {
	DefaultRestHandler
}

// UserResult user information return
type UserResult struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

// Get handle user list api
func (handler *UsersHandler) Get(w http.ResponseWriter, r *http.Request, params map[string]string) {

	users, err := model.GetUsers(-1)
	if err != nil {
		v := fmt.Sprintf("error %v", err)
		w.Write([]byte(v))
		return
	}

	results := make([]*UserResult, 0)
	for _, user := range users {
		result := &UserResult{
			UserID: user.UserID,
			Name:   user.Name,
			Type:   user.Type,
		}
		results = append(results, result)
	}
	values, _ := json.Marshal(results)
	w.Write(values)
}

const userInputSchemaJSON = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"properties": {
		"name": {
			"description": "Name of the user",
			"type": "string",
			"minLength": 1,
			"maxLength": 64
		}
	},
	"required": ["name"]
    
}
`

var inputSchema *jsonschema.Schema

func init() {
	userInputURL := "user_schema.json"
	err := compiler.AddResource(userInputURL, strings.NewReader(userInputSchemaJSON))
	if err != nil {
		log.Printf("error in load user schema %v", err)
		panic("wrong load schema")
	}

	inputSchema, err = compiler.Compile(userInputURL)
	if err != nil {
		log.Printf("error in compile user schema %v", err)
		panic("error in compile user schema")
	}
}

// Post handler user like , unlike update
func (handler *UsersHandler) Post(w http.ResponseWriter, r *http.Request, params map[string]string, body []byte) {
	if err := inputSchema.Validate(bytes.NewReader(body)); err != nil {
		v := fmt.Sprintf("error %v", err)
		w.Write([]byte(v))
		return
	}
	var user UserResult
	if err := json.Unmarshal(body, &user); err != nil {
		v := fmt.Sprintf("error in parse body: %v", err)
		w.Write([]byte(v))
		return
	}

	userUUID, _ := uuid.NewRandom()
	userID := userUUID.Format(uuid.StyleWithoutDash)
	t := "user"
	err := model.CreateNewUser(userID, user.Name, t)
	if err != nil {
		v := fmt.Sprintf("error %v", err)
		w.Write([]byte(v))
		return
	}

	u, err := model.GetUserByID(userID)
	if err != nil {
		v := fmt.Sprintf("error %v", err)
		w.Write([]byte(v))
		return
	}
	result := &UserResult{
		UserID: u.UserID,
		Name:   u.Name,
		Type:   u.Type,
	}

	values, _ := json.Marshal(result)
	w.Write(values)
}
