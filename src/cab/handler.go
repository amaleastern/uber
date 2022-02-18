package cab

import (
	"github.com/gofiber/fiber/v2"
)

// Handler interface
type Handler interface {
	GetNearbyCabs(c *fiber.Ctx) error
}

// HandlerImpl struct
type HandlerImpl struct {
	Service Service
}

// NewHandler func
func NewHandler(service Service) Handler {
	return &HandlerImpl{Service: service}
}

// GetNearbyCabs cab handler
func (h *HandlerImpl) GetNearbyCabs(c *fiber.Ctx) error {

	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	if latitude == "" || longitude == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "latitude and longitude are required"})
	}

	cabs, err := h.Service.GetNearbyCabs(latitude, longitude)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(cabs)
}
