//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"microApiBasic/internal/handler"
	"microApiBasic/internal/repository"
	"microApiBasic/internal/server"
	"microApiBasic/internal/service"
	"microApiBasic/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
