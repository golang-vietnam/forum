package main

import (
	"fmt"
	"github.com/golang-vietnam/forum/cmd"
	"github.com/spf13/viper"
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
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch arg {
	case "start":
		viper.Set("env", "development")
	case "deploy":
		viper.Set("env", "production")
	case "test":
		viper.Set("env", "testing")
	default:
		fmt.Println("Invalid command:", arg)
		showUsage()
		os.Exit(1)
	}
	cmd.Server()
}
