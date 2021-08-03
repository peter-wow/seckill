// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/peter-wow/seckill/app/order/service/internal/biz"
	"github.com/peter-wow/seckill/app/order/service/internal/conf"
	"github.com/peter-wow/seckill/app/order/service/internal/data"
	"github.com/peter-wow/seckill/app/order/service/internal/server"
	"github.com/peter-wow/seckill/app/order/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	registry := server.NewRegistrar()
	dataData, cleanup, err := data.NewData(confData, logger, registry)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUsecase := biz.NewOrderUsecase(orderRepo, logger)
	seckillOrderRepo := data.NewSeckillOrderRepo(dataData, logger)
	seckillOrderUsecase := biz.NewSeckillOrderUsecase(seckillOrderRepo, logger)
	seckillGoodsRepo := data.NewSeckillGoodsRepo(dataData, logger)
	seckillGoodsUsecase := biz.NewSeckillGoodsUsecase(seckillGoodsRepo, logger)
	orderService := service.NewOrderService(orderUsecase, seckillOrderUsecase, seckillGoodsUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, orderService, logger)
	grpcServer := server.NewGRPCServer(confServer, orderService, logger)
	app := newApp(logger, httpServer, grpcServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
