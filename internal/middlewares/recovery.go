package middlewares

import (
	"github.com/gofiber/fiber/v2"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func Recovery() fiber.Handler {
	return fiberrecover.New()
}
