package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viettrung2103/bookmark-management-lesson/docs"
	"github.com/viettrung2103/bookmark-management-lesson/internal/repository"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/utils"

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
	app   *gin.Engine
	cfg   *Config
	redis *redis.Client
}

// NewEngine creates a new engine
func NewEngine(cfg *Config, redis *redis.Client) Engine {
	app := &engine{
		app:   gin.Default(),
		cfg:   cfg,
		redis: redis,
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

	keyGen := utils.NewKeyGenerator()
	shortenUrlRepo := repository.NewUrlStorage(e.redis)
	shortenUrlSvc := service.NewShortenUrl(shortenUrlRepo, keyGen)
	shortenUrlHandler := handler.NewShortenUrlHandler(shortenUrlSvc)

	healthCheckRepo := repository.NewHealthCheck(e.redis)
	healthCheckSvc := service.NewHealthCheck(healthCheckRepo)

	healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)

	e.app.GET("/genpass", genPassHandler.GeneratePassword)
	e.app.GET("/health-check", healthCheckHandler.CheckHealth)

	//int swagger routes
	docs.SwaggerInfo.Host = e.cfg.Hostname
	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	e.app.POST("/v1/links/shorten", shortenUrlHandler.ShortenUrl)
	e.app.GET("/v1/links/shorten/:code", shortenUrlHandler.Redirect)
}
