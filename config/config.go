package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	EnvTesting     = "testing"
	EnvDevelopment = "development"
	EnvProduction  = "production"
	SecretKey      = "secret"
	AvatarPath     = "uploads/avatar/"
)

type database struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}
type server struct {
	Host string
	Port int
}
type envValue struct {
	Database database
	Server   server
}
type env struct {
	Testing     envValue
	Development envValue
	Production  envValue
}
type config struct {
	Env    env
	Secret string
}

type Config config
type Server server
type Env env
type EnvValue envValue
type Database database

var configValue *config

func Loads(filePath string) *config {
	var fileName string
	var yamlFile []byte
	var err error

	if fileName, err = filepath.Abs(filePath); err != nil {
		panic(err)
	}

	if yamlFile, err = ioutil.ReadFile(fileName); err != nil {
		panic(err)
	}
	configValue = &config{}
	if err = yaml.Unmarshal(yamlFile, configValue); err != nil {
		panic(err)
	}
	return configValue
}

func SetEnv(env string) {
	if env != EnvDevelopment && env != EnvProduction && env != EnvTesting {
		panic("Invalid env")
	}
	os.Setenv("env", env)
}

func GetEnv() string {
	env := os.Getenv("env")
	if env == "" {
		panic("Environment not set")
	}
	return env
}
func GetEnvValue() *envValue {
	if configValue == nil {
		panic("Must run Loads first")
	}
	env := GetEnv()
	switch env {
	case EnvDevelopment:
		return &configValue.Env.Development
	case EnvProduction:
		return &configValue.Env.Production
	case EnvTesting:
		return &configValue.Env.Testing
	default:
		return nil
	}
}
func GetSecret() string {
	if configValue == nil {
		panic("Must run Loads first")
	}
	return configValue.Secret
}
