package main

import (
	"fmt"
	"github.com/babon21/redis-impl/internal/app/client/config"
	cacheHttp "github.com/babon21/redis-impl/internal/app/client/delivery/http"
	"github.com/babon21/redis-impl/internal/app/client/gateway"
	"github.com/babon21/redis-impl/internal/app/client/usecase"
	"github.com/babon21/redis-impl/internal/pkg/http/middleware"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.Init()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.AccessLogMiddleware)
	fmt.Println(conf.Server.ServerUrl)
	redisGateway := gateway.NewRedisGateway(conf.Server.ServerUrl)
	redisUsecase := usecase.NewRedisUsecase(redisGateway)
	cacheHttp.NewCacheHandler(e, redisUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
