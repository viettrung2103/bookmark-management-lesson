package main

import (
	"github.com/rs/zerolog/log"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/logger"
)

func main() {

	//ctx := context.Background()
	//
	////redisClient, err := redis.NewClient("")
	////if err != nil {
	////	panic(err)
	////}
	////redisClient.Set(ctx, "key", "1234", time.Hour)
	//
	//urlStorage, err := redis.NewClient("")
	//if err != nil {
	//	panic(err)
	//}
	////urlStorage.Set(ctx, "key", "1234", time.Hour)
	//
	////cacheDB, err := redis.NewClient("CACHE")
	////if err != nil {
	////	panic(err)
	////}
	////cacheDB.Set(ctx, "key", "4567", time.Hour)
	//
	//urlRepo := repository.NewUrlStorage(urlStorage)
	////_ = urlRepo.StoreURL(ctx, "1234", "youtube.com")
	//
	//urlService := service.NewShortenUrl(urlRepo)
	//
	//key, _ := urlService.ShortenUrl(ctx, "https://google.com")
	//print(key)

	logger.SetLogLevel()

	log.Debug().Str("name", "debug").Int("run-time", 1000).Msg("log nay chi hien thi o debug level")
	log.Info().Str("name", "debug").Int("run-time", 1000).Msg("log nay chi hien thi o debug level")

}
