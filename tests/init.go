package tests

import (
	"encoding/json"
	"fmt"
	"github.com/golang-vietnam/forum/cmd"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Error struct {
	Message string `json:message`
	Id      string `json:id`
}
type userModel models.User

const (
	UserColName = models.UserColName
)

var (
	server   string
	userApi  string
	authApi  string
	errorApi string
	serv     chan bool
)

func getConnectString(host string, port int) string {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", host, port))
	if err != nil {
		panic("Url config invalid")
	}
	return u.String()
}

func runServer() {
	cmd.Start()
	serv <- true
}

func init() {
	config.Loads("../config/config.yml")
	config.SetEnv(config.EnvTesting)
	env := config.GetEnvValue()
	server = getConnectString(env.Server.Host, env.Server.Port)
	userApi = server + "/v1/user/"
	authApi = server + "/v1/auth/"
	errorApi = server + "/v1/errors/"
	go runServer()
	select {
	case <-serv:
	case <-time.After(time.Second * 1):
	}
}

type data struct {
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
