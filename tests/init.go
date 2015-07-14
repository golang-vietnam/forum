package tests

import (
	"encoding/json"
	"fmt"
	"github.com/golang-vietnam/forum/cmd"
	"github.com/golang-vietnam/forum/helpers/config"
	"github.com/golang-vietnam/forum/helpers/database"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
type Categories struct {
	Categories []models.Category `json:categories`
}
type userModel models.User
type categoryModel models.Category

var (
	server      string
	userApi     string
	authApi     string
	errorApi    string
	categoryApi string
	serv        chan bool
)

func newObjectId() bson.ObjectId {
	return bson.NewObjectId()
}
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
	config.Loads("../config.yml")
	config.SetEnv(config.EnvTesting)
	env := config.GetEnvValue()
	server = getConnectString(env.Server.Host, env.Server.Port)
	userApi = server + "/v1/users/"
	authApi = server + "/v1/auths/"
	errorApi = server + "/v1/errors/"
	categoryApi = server + "/v1/categories/"
	go runServer()
	select {
	case <-serv:
	case <-time.After(time.Second * 1):
	}
}

func clearAll() {
	database.ClearAll()
}

func userCollection() *mgo.Collection {
	return database.Collection(models.UserColName)
}
func categoryCollection() *mgo.Collection {
	return database.Collection(models.CategoryColName)
}

// method string, urlStr string, model interface{}
func do_request(params ...interface{}) *http.Response {

	if params == nil {
		panic("do_request func must have parameter")
	}

	if len(params) < 2 {
		panic("do_request func must >= 2 parameter")
	}
	method := params[0].(string)
	urlStr := params[1].(string)

	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		panic("Method invalid")
	}
	var data *strings.Reader = strings.NewReader("")

	if len(params) > 2 {
		model := params[2]
		if model != nil {
			jsondata, _ := json.Marshal(model)
			data = strings.NewReader(string(jsondata))
		}
	}
	req, err := http.NewRequest(method, urlStr, data)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	if len(params) > 3 {
		if headers, ok := params[3].(map[string]string); ok {
			for k, v := range headers {
				req.Header.Set(k, v)
			}
		} else {
			panic("hearders must map[string]string type")
		}
	}

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
