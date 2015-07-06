package main

import (
	"fmt"
	"github.com/golang-vietnam/forum/cmd"
	"github.com/golang-vietnam/forum/config"
	"os"
)

func showUsage() {
	fmt.Println("Usage: forum <command>")
	fmt.Println("=============================\n")
	fmt.Println("Avaialable commands:")
	fmt.Println("'forum' or 'forum start'    # run server on develop mode")
	fmt.Println("'forum test' # run server on test mode")
	fmt.Println("'forum deploy' # run server on production mode")
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 1 {
		fmt.Println("Invalid command usage\n")
		showUsage()
		os.Exit(1)
	}
	arg := "start"
	if len(os.Args) == 2 {
		arg = os.Args[1]
	}
	config.Loads("config/config.yml")
	switch arg {
	case "start":
		config.SetEnv(config.EnvDevelopment)
	case "deploy":
		config.SetEnv(config.EnvProduction)
	case "test":
		config.SetEnv(config.EnvTesting)
	default:
		fmt.Println("Invalid command:", arg)
		showUsage()
		os.Exit(1)
	}
	cmd.Start()
}
