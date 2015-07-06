package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/handlers"
	"github.com/golang-vietnam/forum/middleware"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var (
	srv *graceful.Server
)

type server struct {
	srv  *graceful.Server
	done chan bool
}

func Start() error {
	app := setup()
	env := config.GetEnvValue()
	srv = &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    env.Server.Host + ":" + strconv.Itoa(env.Server.Port),
			Handler: app,
		},
	}
	return srv.ListenAndServe()
}

func Stop() {
	if srv == nil {
		panic("Server not running")
	}
	srv.Stop(0)
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
	app.Static("/public", "./public")
	//Set up api v1
	routeV1(app)
	return app
}

func routeV1(app *gin.Engine) {
	//Home
	homeHandler := handlers.NewHomeHandler()
	v1Group := app.Group("/v1")
	{
		v1Group.GET("/", homeHandler.Index)
	}
	apiErrorHandler := handlers.NewErrorHandler()
	apiErrorGroup := v1Group.Group("/errors")
	{
		apiErrorGroup.GET("/", apiErrorHandler.List)
		apiErrorGroup.GET("/:errorId", apiErrorHandler.GetById)
	}
	//User
	userHandler := handlers.NewUserHandler()
	list := []gin.HandlerFunc{userHandler.Create}
	userGroup := v1Group.Group("/user")
	{
		userGroup.GET("/:userId", userHandler.Detail)
		userGroup.PUT("/:userId", userHandler.Edit)
		userGroup.POST("/", list...)
	}

	//Post
	postHandler := handlers.NewPostHandler()
	postGroup := v1Group.Group("/post")
	{
		postGroup.GET("/", postHandler.Index)
		postGroup.POST("/", postHandler.Create)
		postGroup.GET("/:id", postHandler.GetById)
	}

	//Auth
	authHandler := handlers.NewAuthHandler()
	authGroup := v1Group.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}
}
