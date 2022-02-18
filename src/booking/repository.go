package booking

import (
	"uber/src/models"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	GetBooking(bookingID int) (*models.Booking, error)
	GetBookingsByUserID(userID int) ([]models.Booking, error)
	CreateBooking(booking *models.Booking) error
	CancelBookingUser(bookingID int) error
	CancelBookingDriver(bookingID int) error
	AcceptBooking(bookingID int) error
	MarkBookingAsCompleted(bookingID int) error
	MarkCabAsActive(cabID int) error
	MarkCabAsBusy(cabID int) error
}

// RepositoryImpl struct
type RepositoryImpl struct {
	DB *gorm.DB
}

// NewRepository func
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

// CreateBooking Create booking
func (r *RepositoryImpl) CreateBooking(booking *models.Booking) error {
	return r.DB.Create(&booking).Error
}

// MarkBookingAsCompleted Mark booking as completed
func (r *RepositoryImpl) MarkBookingAsCompleted(bookingID int) error {
	booking := models.Booking{ID: bookingID}
	return r.DB.Model(booking).Update("completed", true).Error
}

// MarkCabAsActive Mark cab as active
func (r *RepositoryImpl) MarkCabAsActive(cabID int) error {
	cab := models.Cab{ID: cabID}
	return r.DB.Model(cab).Update("active", true).Error
}

// MarkCabAsActive Mark cab as busy
func (r *RepositoryImpl) MarkCabAsBusy(cabID int) error {
	cab := models.Cab{ID: cabID}
	return r.DB.Model(cab).Update("active", false).Error
}

// CancelBookingDriver mark booking as cancelled by driver
func (r *RepositoryImpl) CancelBookingDriver(bookingID int) error {
	booking := models.Booking{ID: bookingID}
	return r.DB.Model(booking).Update("cancelled_by_driver", true).Error
}

// CancelBookingUser mark booking as cancelled by user
func (r *RepositoryImpl) CancelBookingUser(bookingID int) error {
	booking := models.Booking{ID: bookingID}
	return r.DB.Model(booking).Update("cancelled_by_user", true).Error
}

// AcceptBooking mark booking as accepeted
func (r *RepositoryImpl) AcceptBooking(bookingID int) error {
	booking := models.Booking{ID: bookingID}
	return r.DB.Model(booking).Update("accepted_by_driver", true).Error
}

// GetBookingsByUserID Get bookings by id_user
func (r *RepositoryImpl) GetBookingsByUserID(userID int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.Where("id_user = ?", userID).Find(&bookings).Error
	return bookings, err
}

// GetBooking Get booking by ibooking id
func (r *RepositoryImpl) GetBooking(bookingID int) (*models.Booking, error) {
	var booking models.Booking
	err := r.DB.Where("id = ?", bookingID).First(&booking).Error
	return &booking, err
}
