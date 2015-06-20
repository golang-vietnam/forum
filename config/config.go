package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

const (
	ENV             = "env"
	ENV_TESTING     = "testing"
	ENV_DEVELOPMENT = "development"
	ENV_PRODUCTION  = "production"
	CONFIG_NAME     = "config"
	CONFIG_PATH     = "config"
)

func mapDb(confEnv map[string]interface{}) map[string]interface{} {
	database := confEnv["database"]
	databaseMap, _ := database.(map[string]interface{})
	return databaseMap
}
func SetEnv(env string) string {
	configPath := CONFIG_PATH
	viper.SetConfigName(CONFIG_NAME)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.Set(ENV, env)
	u, err := url.Parse(fmt.Sprintf("http://%s:%s", GetServer("host"), GetServer("port")))
	if err != nil {
		panic("Url config invalid")
	}
	return u.String()
}

func GetEnv() interface{} {
	return viper.Get(ENV)
}

func GetDB(key string) string {
	env := viper.Get(ENV)
	confEnv := viper.GetStringMap(ENV_DEVELOPMENT)
	switch env {
	case ENV_DEVELOPMENT:
		value, _ := mapDb(confEnv)[key].(string)
		return value
	case ENV_PRODUCTION:
		confEnv := viper.GetStringMap(ENV_PRODUCTION)
		value, _ := mapDb(confEnv)[key].(string)
		return value
	case ENV_TESTING:
		confEnv := viper.GetStringMap(ENV_TESTING)
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
	env := viper.Get(ENV)
	confEnv := viper.GetStringMap(ENV_DEVELOPMENT)
	switch env {
	case ENV_DEVELOPMENT:
		value, _ := mapServer(confEnv)[key].(string)
		return value
	case ENV_PRODUCTION:
		confEnv := viper.GetStringMap(ENV_PRODUCTION)
		value, _ := mapServer(confEnv)[key].(string)
		return value
	case ENV_TESTING:
		confEnv := viper.GetStringMap(ENV_TESTING)
		value, _ := mapServer(confEnv)[key].(string)
		return value
	default:
		log.Panic("Ohshit! Not config environment")
	}
	return ""
}
