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
