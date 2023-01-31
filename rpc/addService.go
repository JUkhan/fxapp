package rpc

import (
	"context"
	"fxapp/proto"

	"go.uber.org/fx"
)

type AddService = proto.AddServiceServer

type addService struct {
	proto.UnimplementedAddServiceServer
}

func newService() (AddService, error) {
	return &addService{}, nil
}

func (s *addService) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *addService) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

// Module provided to fx
var AddServiceModule = fx.Options(
	fx.Provide(newService),
)
