package booking

import (
	"errors"
	"strconv"
	"time"
	"uber/src/models"
)

// Service interface
type Service interface {
	GetBookingsByUser(userID string) ([]models.Booking, error)
	CreateBooking(booking *models.Booking) error
	CancelBooking(bookingID string) error
	AcceptBooking(bookingID string) error
	MarkRideAsCompleted(bookingID string) error
}

// ServiceImpl struct
type ServiceImpl struct {
	Repository Repository
}

// NewService func
func NewService(repository Repository) Service {
	return &ServiceImpl{Repository: repository}
}

// CreateBooking Create booking
func (s *ServiceImpl) CreateBooking(booking *models.Booking) error {
	booking.BookingDate = time.Now()
	return s.Repository.CreateBooking(booking)
}

// GetBookingsByUser Get bookings by user
func (s *ServiceImpl) GetBookingsByUser(userID string) ([]models.Booking, error) {

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	return s.Repository.GetBookingsByUserID(userIDInt)
}

// AcceptBooking Accept booking
func (s *ServiceImpl) AcceptBooking(bookingID string) error {
	bookingIDInt, err := strconv.Atoi(bookingID)
	if err != nil {
		return err
	}

	booking, err := s.checkBooking(bookingIDInt, "accept")
	if err != nil {
		return err
	}

	if err = s.Repository.MarkCabAsBusy(booking.IDCab); err != nil {
		return err
	}

	return s.Repository.AcceptBooking(bookingIDInt)
}

// MarkRideAsCompleted Mark ride as completed
func (s *ServiceImpl) MarkRideAsCompleted(bookingID string) error {
	bookingIDInt, err := strconv.Atoi(bookingID)
	if err != nil {
		return err
	}

	booking, err := s.checkBooking(bookingIDInt, "complete")
	if err != nil {
		return err
	}

	if err = s.Repository.MarkBookingAsCompleted(bookingIDInt); err != nil {
		return err
	}

	return s.Repository.MarkCabAsActive(booking.IDCab)
}

// CancelBooking Cancel booking
func (s *ServiceImpl) CancelBooking(bookingID string) error {
	bookingIDInt, err := strconv.Atoi(bookingID)
	if err != nil {
		return err
	}

	if _, err = s.checkBooking(bookingIDInt, "cancel"); err != nil {
		return err
	}

	return s.Repository.CancelBookingUser(bookingIDInt) //check role (user/driver) by auth token
	// return s.Repository.CancelBookingDriver(bookingIDInt)
}

func (s *ServiceImpl) checkBooking(bookingID int, context string) (*models.Booking, error) {
	booking, err := s.Repository.GetBooking(bookingID)
	if err != nil {
		return nil, err
	}

	if booking.CancelledByDriver {
		return nil, errors.New("cannot " + context + " booking : already cancelled by driver")
	} else if booking.CancelledByUser {
		return nil, errors.New("cannot " + context + " booking : already cancelled by user")
	} else if booking.Completed {
		return nil, errors.New("cannot " + context + " booking : already completed")
	}

	if context == "complete" {
		if !booking.AcceptedByDriver {
			return nil, errors.New("cannot " + context + "  booking : not accepeted by driver")
		}
	} else {
		if booking.AcceptedByDriver {
			return nil, errors.New("cannot " + context + " booking : already accepeted by driver")
		}
	}

	return booking, nil
}
