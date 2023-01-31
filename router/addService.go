package router

import (
	"context"
	"fmt"
	"fxapp/proto"
	"fxapp/rpc"
	"fxapp/server"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type addServiceModule struct {
	*server.Server
	client proto.AddServiceClient
}

func newAddServiceModule(server *server.Server, conn *rpc.Connector) *addServiceModule {
	client := proto.NewAddServiceClient(conn.Conn)
	service := addServiceModule{server, client}
	service.setupRoutes()
	return &service
}

func (b *addServiceModule) setupRoutes() {
	b.App.Get("/add/:a/:b", b.add)
	b.App.Get("/mult/:a/:b", b.mul)

}

func (s *addServiceModule) add(c *fiber.Ctx) error {
	a, err := strconv.ParseUint(c.Params("a"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument A",
		})
	}
	b, err := strconv.ParseUint(c.Params("b"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument B",
		})
	}
	req := &proto.Request{A: int64(a), B: int64(b)}
	if res, err := s.client.Add(context.Background(), req); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": fmt.Sprint(res.Result),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})

}

func (s *addServiceModule) mul(c *fiber.Ctx) error {
	a, err := strconv.ParseUint(c.Params("a"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument A",
		})
	}
	b, err := strconv.ParseUint(c.Params("b"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument B",
		})
	}
	req := &proto.Request{A: int64(a), B: int64(b)}
	if res, err := s.client.Multiply(context.Background(), req); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": fmt.Sprint(res.Result),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})

}

var AddServiceModule = fx.Options(
	fx.Invoke(newAddServiceModule),
)
