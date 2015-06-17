package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/database"
	"io/ioutil"
	"net/http"
	"strings"
)

type Error struct {
	Message string `json:message`
	Id      string `json:id`
}

var (
	server  string
	userApi string
	authApi string
)

func init() {
	server = config.SetEnv(config.ENV_TESTING)
	userApi = server + "/v1/user/"
	authApi = server + "/v1/auth/"
	database.InitDb()
}

func do_request(method string, urlStr string, model interface{}) *http.Response {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		panic("Method invalid")
	}
	var data *strings.Reader = nil
	if model != nil {
		jsondata, _ := json.Marshal(model)
		data = strings.NewReader(string(jsondata))
	}
	req, err := http.NewRequest(method, urlStr, data)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

func parse_response(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	return body
}
