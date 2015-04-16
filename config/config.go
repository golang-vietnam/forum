package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func mapDb(confEnv map[string]interface{}) map[string]interface{} {
	database := confEnv["database"]
	databaseMap, _ := database.(map[string]interface{})
	return databaseMap
}
func GetDB(key string) string {
	env := viper.Get("env")
	confEnv := viper.GetStringMap("development")
	switch env {
	case "development":
		value, _ := mapDb(confEnv)[key].(string)
		return value
	case "production":
		confEnv := viper.GetStringMap("production")
		value, _ := mapDb(confEnv)[key].(string)
		return value
	case "testing":
		confEnv := viper.GetStringMap("testing")
		value, _ := mapDb(confEnv)[key].(string)
		return value
	default:
		log.Panic("Ohshit! Not config environment")
	}
	return ""
}
func mapServer(confEnv map[string]interface{}) map[string]interface{} {
	server := confEnv["server"]
	serverMap, _ := server.(map[string]interface{})
	return serverMap
}
func GetServer(key string) string {
	env := viper.Get("env")
	confEnv := viper.GetStringMap("development")
	switch env {
	case "development":
		value, _ := mapServer(confEnv)[key].(string)
		return value
	case "production":
		confEnv := viper.GetStringMap("production")
		value, _ := mapServer(confEnv)[key].(string)
		return value
	case "testing":
		confEnv := viper.GetStringMap("testing")
		value, _ := mapServer(confEnv)[key].(string)
		return value
	default:
		fmt.Println("Ohshit! Not config environment")
	}
	return ""
}
