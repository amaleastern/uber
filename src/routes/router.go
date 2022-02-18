package routes

import (
	"uber/src/healthcheck"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/gorm"
)

// Setup setup routes
func Setup(app *fiber.App, dbHandler *gorm.DB) {

	app.Get("/", healthcheck.Index)

	// for testing purpose
	userHandler := bindUser(dbHandler)
	app.Post("/user", userHandler.AddUser)

	a := app.Group("")

	a.Use(basicauth.New(basicauth.Config{
		Authorizer: Authorizer,
	}))

	// a.Use(Authorizer)

	cabHandler := bindCab(dbHandler)
	a.Get("/cabs", cabHandler.GetNearbyCabs)

	bookingHandler := bindBooking(dbHandler)
	a.Get("/bookings/:UserID", bookingHandler.GetBookingsByUser) // userID can be fetched from auth token
	a.Post("/book", bookingHandler.BookRide)
	a.Put("/booking/:BookingID/cancel", bookingHandler.CancelBooking)
	a.Put("/booking/:BookingID/accept", bookingHandler.AcceptBooking)
	a.Put("/booking/:BookingID/complete", bookingHandler.MarkRideAsCompleted)
}
