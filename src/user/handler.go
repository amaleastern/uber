package user

import (
	"uber/src/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

// Handler interface
type Handler interface {
	AddUser(c *fiber.Ctx) error
}

// HandlerImpl struct
type HandlerImpl struct {
	Service Service
}

// NewHandler func
func NewHandler(service Service) Handler {
	return &HandlerImpl{Service: service}
}

// AddUser user handler
func (h *HandlerImpl) AddUser(c *fiber.Ctx) error {

	var user models.User
	var err error

	if err = c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if _, err = govalidator.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = h.Service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Created"})
}
