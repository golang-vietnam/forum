package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/config"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

func getServer() string {
	viper.Set("env", "testing")
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return "http://" + config.GetServer("host") + ":" + config.GetServer("port")
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
