package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	. "gopkg.in/check.v1"
)

func (t *HandlerTest) getUserRelations(c *C, userID string, code int) []map[string]interface{} {

	relationGetURL := "http://" + t.port + "/users" + "/" + userID + "/relationships"
	resp, err := http.Get(relationGetURL)
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, code)

	var relations = make([]map[string]interface{}, 0)
	if code == 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &relations)
	}
	return relations
}

func (t *HandlerTest) TestRelationHandler(c *C) {

	postUrl := "http://" + t.port + "/users"
	postData := `{"name": "liuzhenzhong"}`
	http.Post(postUrl, "application/json", strings.NewReader(postData))

	postData = `{"name": "liuzhenzhong"}`
	http.Post(postUrl, "application/json", strings.NewReader(postData))

	getUrl := "http://" + t.port + "/users"
	resp, _ := http.Get(getUrl)
	data, _ := ioutil.ReadAll(resp.Body)
	var users = make([]map[string]interface{}, 0)
	json.Unmarshal(data, &users)

	c.Assert(len(users), Equals, 2)

	t.getUserRelations(c, "23399993", 404)

	userID1 := users[0]["id"].(string)
	userID2 := users[1]["id"].(string)

	var relations = t.getUserRelations(c, userID1, 200)
	c.Assert(len(relations), Equals, 0)

	t.putRelation(c, userID1, userID2, "like", 400)

	t.putRelation(c, userID1, userID2, "liked", 200)
	t.putRelation(c, userID2, userID1, "liked", 200)

	// error for resend logic
	t.putRelation(c, userID2, userID1, "liked", 500)

	relations = t.getUserRelations(c, userID1, 200)
	c.Assert(len(relations), Equals, 1)
	c.Assert(relations[0]["state"], Equals, "matched")

	relations = t.getUserRelations(c, userID2, 200)
	c.Assert(len(relations), Equals, 1)
	c.Assert(relations[0]["state"], Equals, "matched")
}

func (t *HandlerTest) putRelation(c *C, userID1, userID2 string, state string, code int) {
	putRelationPost := "http://" + t.port + "/users" + "/" + userID1 + "/relationships" + "/" + userID2
	client := &http.Client{}
	putData := fmt.Sprintf(`{"state": "%v"}`, state)
	req, err := http.NewRequest("PUT", putRelationPost, strings.NewReader(putData))
	resp, err := client.Do(req)
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, code)
}
