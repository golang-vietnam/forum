package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/config"
	"github.com/golang-vietnam/forum/helpers/database"
	"github.com/golang-vietnam/forum/middleware"
	"runtime"
	"strconv"
)

func Start() error {
	app := setup()
	env := config.GetEnvValue()
	return app.Run(env.Server.Host + ":" + strconv.Itoa(env.Server.Port))
}

func setup() *gin.Engine {

	runtime.GOMAXPROCS(runtime.NumCPU())
	if _, err := database.InitDb(); err != nil {
		panic(err)
	}
	app := gin.New()

	app.Use(func(c *gin.Context) {
		c.Set(config.SecretKey, config.GetSecret())
		c.Next()
	})

	if config.GetEnv() == config.EnvProduction {
		app.Use(middleware.Recovery())
	} else {
		app.Use(gin.Recovery())
	}

	if config.GetEnv() == config.EnvTesting {
		gin.SetMode(gin.TestMode)
	} else {
		app.Use(gin.Logger())
	}

	app.Use(middleware.ErrorHandler())
	//Set up api v1
	routeV1(app)
	return app
}
