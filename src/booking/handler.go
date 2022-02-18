package booking

import (
	"uber/src/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

// Handler interface
type Handler interface {
	GetBookingsByUser(c *fiber.Ctx) error
	BookRide(c *fiber.Ctx) error
	CancelBooking(c *fiber.Ctx) error
	AcceptBooking(c *fiber.Ctx) error
	MarkRideAsCompleted(c *fiber.Ctx) error
}

// HandlerImpl struct
type HandlerImpl struct {
	Service Service
}

// NewHandler func
func NewHandler(service Service) Handler {
	return &HandlerImpl{Service: service}
}

// GetBookingsByUser booking handler
func (h *HandlerImpl) GetBookingsByUser(c *fiber.Ctx) error {
	// userID := c.Context().UserValue("UserID").(int)

	userID := c.Params("UserID")

	cabs, err := h.Service.GetBookingsByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(cabs)
}

// BookRide booking handler
func (h *HandlerImpl) BookRide(c *fiber.Ctx) error {
	var booking models.Booking
	var err error

	if err = c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if _, err = govalidator.ValidateStruct(booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = h.Service.CreateBooking(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Booked"})
}

// CancelRide booking handler
func (h *HandlerImpl) CancelBooking(c *fiber.Ctx) error {
	bookingID := c.Params("BookingID")

	if err := h.Service.CancelBooking(bookingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Updated"})
}

// AcceptBooking booking handler
func (h *HandlerImpl) AcceptBooking(c *fiber.Ctx) error {
	bookingID := c.Params("BookingID")

	if err := h.Service.AcceptBooking(bookingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Updated"})
}

// MarkRideAsCompleted booking handler
func (h *HandlerImpl) MarkRideAsCompleted(c *fiber.Ctx) error {
	bookingID := c.Params("BookingID")

	if err := h.Service.MarkRideAsCompleted(bookingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Updated"})
}
