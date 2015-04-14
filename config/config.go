package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func mapDb(confEnv map[string]interface{}) map[string]interface{} {
	database := confEnv["database"]
	databaseMap, _ := database.(map[string]interface{})
	return databaseMap
}
func GetDBHost() string {
	env := viper.Get("env")
	confEnv := viper.GetStringMap("development")
	switch env {
	case "development":
		host, _ := mapDb(confEnv)["host"].(string)
		return host
	case "production":
		confEnv := viper.GetStringMap("production")
		host, _ := mapDb(confEnv)["host"].(string)
		return host
	case "testing":
		confEnv := viper.GetStringMap("testing")
		host, _ := mapDb(confEnv)["host"].(string)
		return host
	default:
		fmt.Println("Ohshit! Not config environment")
	}
	return ""
}
