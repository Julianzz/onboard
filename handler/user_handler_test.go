package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	. "gopkg.in/check.v1"
)

func (t *HandlerTest) checkUsers(c *C) []map[string]interface{} {
	getUrl := "http://" + t.port + "/users"
	resp, err := http.Get(getUrl)
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	data, err := ioutil.ReadAll(resp.Body)
	c.Assert(err, IsNil)
	fmt.Println(string(data))
	var values = make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &values)
	c.Assert(err, IsNil)
	return values
}

func (t *HandlerTest) TestUserHandler(c *C) {

	users := t.checkUsers(c)
	c.Assert(len(users), Equals, 0)

	postUrl := "http://" + t.port + "/users"
	postData := `{"name": "liuzhenzhong"`
	resp, err := http.Post(postUrl, "application/json", strings.NewReader(postData))
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 400)

	postData = `{"name": "liuzhenzhong"}`
	resp, err = http.Post(postUrl, "application/json", strings.NewReader(postData))
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	users = t.checkUsers(c)
	c.Assert(len(users), Equals, 1)

	postUrl = "http://" + t.port + "/users"
	postData = `{"name": "liuzhenzhong"}`
	resp, err = http.Post(postUrl, "application/json", strings.NewReader(postData))
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	users = t.checkUsers(c)
	c.Assert(len(users), Equals, 2)
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

	relationPost := "http://" + t.port + "/users" + "/23399993" + "/relationships"
	resp, err := http.Get(relationPost)
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 404)

	userID1 := users[0]["id"].(string)
	userID2 := users[1]["id"].(string)

	relationPost = "http://" + t.port + "/users" + "/" + userID1 + "/relationships"
	resp, err = http.Get(relationPost)
	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	data, _ = ioutil.ReadAll(resp.Body)
	var relations = make([]map[string]interface{}, 0)
	json.Unmarshal(data, &relations)
	c.Assert(len(relations), Equals, 0)

	t.putRelation(c, userID1, userID2, "like", 400)

	t.putRelation(c, userID1, userID2, "liked", 200)
	t.putRelation(c, userID2, userID1, "liked", 200)

	// error for resend logic
	t.putRelation(c, userID2, userID1, "liked", 500)
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
