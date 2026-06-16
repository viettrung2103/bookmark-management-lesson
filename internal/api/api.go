package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viettrung2103/bookmark-management-lesson/docs"
	"github.com/viettrung2103/bookmark-management-lesson/internal/app/handler"
	userHld "github.com/viettrung2103/bookmark-management-lesson/internal/app/handler/user"
	"github.com/viettrung2103/bookmark-management-lesson/internal/app/repository"
	userRepository "github.com/viettrung2103/bookmark-management-lesson/internal/app/repository/user"
	"github.com/viettrung2103/bookmark-management-lesson/internal/app/service"
	userService "github.com/viettrung2103/bookmark-management-lesson/internal/app/service/user"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/utils"
	"gorm.io/gorm"

	_ "github.com/viettrung2103/bookmark-management-lesson/docs"
)

// Engine represents the application engine
type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	initHandlers() *handlers
}

type engine struct {
	app   *gin.Engine
	cfg   *Config
	redis *redis.Client
	db    *gorm.DB
}

// NewEngine creates a new engine
func NewEngine(app *gin.Engine, cfg *Config, redis *redis.Client, db *gorm.DB) Engine {
	eng := &engine{
		//app:   gin.Default(),
		app:   app,
		cfg:   cfg,
		redis: redis,
		db:    db,
	}
	eng.initRoutes()

	return eng
}

// Start starts the engine
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// ServeHTTP implements the http.Handler interface to handle HTTP requests, to serve a specific request for testing purpose
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

type handlers struct {
	healthCheckHandler handler.HealthCheck
	linkHandler        handler.ShortenUrl
	userHandler        userHld.Handler
}

func (e *engine) initHandlers() *handlers {
	healthCheckRepo := repository.NewHealthCheck(e.redis)
	shortenUrlRepo := repository.NewUrlStorage(e.redis)

	keyGen := utils.NewKeyGenerator()

	shortenUrlSvc := service.NewShortenUrl(shortenUrlRepo, keyGen)
	healthCheckSvc := service.NewHealthCheck(healthCheckRepo)

	userRepo := userRepository.NewRepository(e.db)
	userSvc := userService.NewService(userRepo)
	userHanlder := userHld.NewHandler(userSvc)

	return &handlers{
		healthCheckHandler: handler.NewHealthCheck(healthCheckSvc),
		linkHandler:        handler.NewShortenUrlHandler(shortenUrlSvc),
		userHandler:        userHanlder,
	}

}

// initRoutes initializes the routes
func (e *engine) initRoutes() {
	// genpass svc, handle and route
	//genPassSvc := service.NewGenPass()
	//genPassHandler := handler.NewGenPass(genPassSvc)
	//
	//keyGen := utils.NewKeyGenerator()
	//shortenUrlRepo := repository.NewUrlStorage(e.redis)
	//shortenUrlSvc := service.NewShortenUrl(shortenUrlRepo, keyGen)
	//shortenUrlHandler := handler.NewShortenUrlHandler(shortenUrlSvc)
	//
	//healthCheckRepo := repository.NewHealthCheck(e.redis)
	//healthCheckSvc := service.NewHealthCheck(healthCheckRepo)
	//
	//healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)
	allHandlers := e.initHandlers()

	//e.app.GET("/genpass", genPassHandler.GeneratePassword)
	//e.app.GET("/health-check", healthCheckHandler.CheckHealth)
	//
	//e.app.POST("/v1/links/shorten", shortenUrlHandler.ShortenUrl)
	//e.app.GET("/v1/links/redirect/:code", shortenUrlHandler.Redirect)

	//health-check
	e.app.GET("/health-check", allHandlers.healthCheckHandler.CheckHealth)
	//e.app.GET("/genpass", allHandlers.genPassHandler.GeneratePassword)

	//Init swagger routes
	docs.SwaggerInfo.Host = e.cfg.Hostname
	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Routes := e.app.Group("/v1")
	{ //link routes
		v1Routes.POST("/links/shorten", allHandlers.linkHandler.ShortenUrl)
		v1Routes.GET("/links/redirect/:code", allHandlers.linkHandler.Redirect)

		//user routes
		v1Routes.POST("/users/register", allHandlers.userHandler.Register)
	}

}
