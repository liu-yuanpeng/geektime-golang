// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"week13/app/interface/internal/biz"
	"week13/app/interface/internal/conf"
	"week13/app/interface/internal/data"
	"week13/app/interface/internal/server"
	"week13/app/interface/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery)
	dataData := data.NewData(confData, userClient)
	userRepo := data.NewUserRepo(dataData, auth)
	authUseCase := biz.NewAuthUseCase(auth)
	userUseCase := biz.NewUserUseCase(userRepo, authUseCase)
	interfaceService := service.NewInterfaceService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, interfaceService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
	}, nil
}
