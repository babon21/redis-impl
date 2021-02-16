package main

import (
	"github.com/babon21/redis-impl/internal/app/server/config"
	cacheHttp "github.com/babon21/redis-impl/internal/app/server/delivery/http"
	"github.com/babon21/redis-impl/internal/app/server/repository"
	"github.com/babon21/redis-impl/internal/app/server/usecase"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.Init()

	e := echo.New()

	redisStore := repository.NewInMemoryRedisStore()
	redisUsecase := usecase.NewRedisUsecase(redisStore)
	cacheHttp.NewCacheHandler(e, redisUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
