package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

// Index healthcheck endpoint
func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "OK"})
}
