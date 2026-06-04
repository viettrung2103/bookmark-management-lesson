package main

import "github.com/viettrung2103/bookmark-management-lesson/internal/api"

// @title Bookmark API
// @version 1.0
// @description API for bookmark management
// @host localhost:8080
// @BasePath /
func main() {
	//create app config
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
