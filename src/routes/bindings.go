package routes

import (
	"uber/src/booking"
	"uber/src/cab"
	"uber/src/user"

	"gorm.io/gorm"
)

var userRepository user.Repository

func bindCab(dbHandler *gorm.DB) cab.Handler {
	repository := cab.NewRepository(dbHandler)
	service := cab.NewService(repository)
	handler := cab.NewHandler(service)
	return handler
}

func bindBooking(dbHandler *gorm.DB) booking.Handler {
	repository := booking.NewRepository(dbHandler)
	service := booking.NewService(repository)
	handler := booking.NewHandler(service)
	return handler
}

func bindUser(dbHandler *gorm.DB) user.Handler {
	userRepository = user.NewRepository(dbHandler)
	service := user.NewService(userRepository)
	handler := user.NewHandler(service)
	return handler
}
