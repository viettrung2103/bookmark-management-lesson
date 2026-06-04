package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/viettrung2103/bookmark-management-lesson/docs"
	"github.com/viettrung2103/bookmark-management-lesson/internal/handler"
	"github.com/viettrung2103/bookmark-management-lesson/internal/service"
)

// Engine represents the application engine
type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

// NewEngine creates a new engine
func NewEngine(cfg *Config) Engine {
	app := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	app.initRoutes()

	return app
}

// Start starts the engine
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// ServeHTTP implements the http.Handler interface to handle HTTP requests, to serve a specific request for testing purpose
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

// initRoutes initializes the routes
func (e *engine) initRoutes() {
	// genpass svc, handle and route
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	e.app.GET("/genpass", genPassHandler.GeneratePassword)
	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
