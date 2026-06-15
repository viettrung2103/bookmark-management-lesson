package main

import (
	"github.com/viettrung2103/bookmark-management-lesson/internal/api"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/logger"
	pkgRedis "github.com/viettrung2103/bookmark-management-lesson/pkg/redis"
)

// @title Bookmark API
// @version 2.0
// @description API for bookmark management
// @host localhost:8080
// @BasePath /
func main() {
	logger.SetLogLevel()
	//create app config
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	redis, err := pkgRedis.NewClient("")
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg, redis)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
