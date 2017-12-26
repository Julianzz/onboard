package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/golang-plus/uuid"
	"github.com/santhosh-tekuri/jsonschema"

	"github.com/p1cn/onboard/liuzhenzhong/model"
)

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
func (handler *UsersHandler) Get(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {

	users, err := model.GetUsers(-1)
	if err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in fetch users lists")
	}

	// gen users result list
	results := make([]*UserResult, 0)
	for _, user := range users {
		result := &UserResult{
			UserID: user.UserID,
			Name:   user.Name,
			Type:   user.Type,
		}
		results = append(results, result)
	}
	return results, nil
}

// Post handler user like , unlike update
func (handler *UsersHandler) Post(w http.ResponseWriter, r *http.Request, params map[string]string, body []byte) (interface{}, error) {
	if err := inputSchema.Validate(bytes.NewReader(body)); err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in checking int")
	}

	var user UserResult
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in parsing body")
	}

	// gen userid and save it
	userUUID, _ := uuid.NewRandom()
	userID := userUUID.Format(uuid.StyleWithoutDash)
	t := "user"
	err := model.CreateNewUser(userID, user.Name, t)
	if err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in creating new user")
	}

	// fetch user information to return
	u, err := model.GetUserByID(userID)
	if err != nil {
		return nil, NewRestfulError(err, http.StatusBadRequest, "error in find user")
	}
	result := &UserResult{
		UserID: u.UserID,
		Name:   u.Name,
		Type:   u.Type,
	}
	return result, nil
}
