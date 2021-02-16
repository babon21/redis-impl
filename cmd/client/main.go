package main

import (
	"fmt"
	"github.com/babon21/redis-impl/internal/app/client/config"
	cacheHttp "github.com/babon21/redis-impl/internal/app/client/delivery/http"
	"github.com/babon21/redis-impl/internal/app/client/gateway"
	"github.com/babon21/redis-impl/internal/app/client/usecase"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.Init()

	e := echo.New()
	fmt.Println(conf.Server.ServerUrl)
	redisGateway := gateway.NewRedisGateway(conf.Server.ServerUrl)
	redisUsecase := usecase.NewRedisUsecase(redisGateway)
	cacheHttp.NewCacheHandler(e, redisUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
