package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/viettrung2103/bookmark-management-lesson/internal/api"
	"github.com/viettrung2103/bookmark-management-lesson/internal/app/model"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/logger"
	pkgRedis "github.com/viettrung2103/bookmark-management-lesson/pkg/redis"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/sqldb"
	"gorm.io/gorm"
)

// @title Bookmark API
// @version 1.5
// @description API for bookmark management
// @host localhost:8080
// @BasePath /
func main() {
	logger.SetLogLevel()
	//create app config
	//cfg, err := api.NewConfig()
	//if err != nil {
	//	panic(err)
	//}
	//redis, err := pkgRedis.NewClient("")
	//if err != nil {
	//	panic(err)
	//}
	//
	//dbClient, err := sqldb.NewClient("")
	//if err != nil {
	//	panic(err)
	//}
	//
	//app := api.NewEngine(cfg, redis, dbClient)
	//err = app.Start()
	//if err != nil {
	//	panic(err)
	//}

	// init app config
	cfg := createAPIConfig()

	// init redis
	redisClient := createRedisClient()

	//init db
	db := createDBClient()

	//init app
	app := createAPIApp(cfg, redisClient, db)

	//start app
	err := app.Start()
	if err != nil {
		panic(err)
	}

}

func createDBClient() *gorm.DB {
	dbClient, err := sqldb.NewClient("")
	if err != nil {
		panic(err)
	}
	dbClient.AutoMigrate(&model.User{})
	return dbClient
}

func createRedisClient() *redis.Client {
	redisClient, err := pkgRedis.NewClient("")
	if err != nil {
		panic(err)
	}
	return redisClient
}

func createAPIConfig() *api.Config {
	apiConfig, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	return apiConfig
}

func createAPIApp(cfg *api.Config, redis *redis.Client, db *gorm.DB) api.Engine {
	app := gin.New()
	a := api.NewEngine(app, cfg, redis, db)

	return a
}
