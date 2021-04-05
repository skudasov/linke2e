package mock_api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	fiblog "github.com/gofiber/fiber/v2/middleware/logger"
)

type MockResponse struct {
	Name string
	Data string
}

type CallResponse struct {
	Times int
}

type StubBody map[string]interface{}

var (
	CALLS = make(map[string]int)
	STUB  map[string]interface{}
)

func Start() {
	app := fiber.New()
	app.Use(fiblog.New())

	app.Post("/set_stub_response", func(c *fiber.Ctx) error {
		var stubBody StubBody
		if err := c.BodyParser(&stubBody); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		STUB = stubBody
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/check_call/:path", func(c *fiber.Ctx) error {
		if _, ok := CALLS[c.Params("path")]; !ok {
			return c.JSON(&CallResponse{
				Times: 0,
			})
		}
		return c.JSON(&CallResponse{
			Times: CALLS[c.Params("path")],
		})
	})

	app.Get("/api_stub", func(c *fiber.Ctx) error {
		pathKey := strings.Replace(c.Path(), "/", "", -1)
		CALLS[pathKey] += 1
		return c.JSON(STUB)
	})

	_ = app.Listen(":9092")
}
