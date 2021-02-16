package main

import (
	cacheHttp "github.com/babon21/redis-impl/internal/app/client/delivery/http"
	"github.com/babon21/redis-impl/internal/app/client/gateway"
	"github.com/babon21/redis-impl/internal/app/client/usecase"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	e := echo.New()

	redisGateway := gateway.NewRedisGateway("http://localhost:8080")
	redisUsecase := usecase.NewRedisUsecase(redisGateway)
	cacheHttp.NewCacheHandler(e, redisUsecase)

	log.Fatal().Msg(e.Start(":" + "8081").Error())
}
