package booking

import (
	"errors"
	"testing"
	"uber/src/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var mockRepository *MockRepository

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetBooking(bookingID int) (*models.Booking, error) {
	args := m.Called(bookingID)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Booking), nil
}

func (m *MockRepository) GetBookingsByUserID(userID int) ([]models.Booking, error) {
	args := m.Called(userID)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Booking), nil
}

func (m *MockRepository) CreateBooking(booking *models.Booking) error {
	args := m.Called(booking)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) CancelBookingUser(bookingID int) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) CancelBookingDriver(bookingID int) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) AcceptBooking(bookingID int) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) MarkBookingAsCompleted(bookingID int) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) MarkCabAsActive(cabID int) error {
	args := m.Called(cabID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockRepository) MarkCabAsBusy(cabID int) error {
	args := m.Called(cabID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

type ServiceSuite struct {
	suite.Suite

	service Service
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	mockRepository = new(MockRepository)
	s.service = NewService(mockRepository)
}

func (s *ServiceSuite) TestCreateBooking() {

	var booking models.Booking

	mockRepository.On("CreateBooking", &booking).Return(nil)

	err := s.service.CreateBooking(&booking)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBooking() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("MarkCabAsBusy", booking.IDCab).Return(nil)
	mockRepository.On("AcceptBooking", bookingID).Return(nil)

	err := s.service.AcceptBooking("1")
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError() {
	err := s.service.AcceptBooking("")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError2() {
	bookingID := 1

	mockRepository.On("GetBooking", bookingID).Return(nil, errors.New("ERROR"))

	err := s.service.AcceptBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError3() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("MarkCabAsBusy", booking.IDCab).Return(errors.New("ERROR"))

	err := s.service.AcceptBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError4() {
	bookingID := 1
	var booking models.Booking
	booking.Completed = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)

	err := s.service.AcceptBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError5() {
	bookingID := 1
	var booking models.Booking
	booking.CancelledByUser = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)

	err := s.service.AcceptBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestAcceptBookingError6() {
	bookingID := 1
	var booking models.Booking
	booking.CancelledByDriver = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)

	err := s.service.AcceptBooking("1")
	assert.Error(s.T(), err)
}
func (s *ServiceSuite) TestMarkRideAsCompleted() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1
	booking.AcceptedByDriver = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("MarkBookingAsCompleted", bookingID).Return(nil)
	mockRepository.On("MarkCabAsActive", booking.IDCab).Return(nil)

	err := s.service.MarkRideAsCompleted("1")
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestMarkRideAsCompletedError() {
	err := s.service.MarkRideAsCompleted("")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestMarkRideAsCompletedError2() {
	bookingID := 1
	var booking models.Booking

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)

	err := s.service.MarkRideAsCompleted("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestMarkRideAsCompletedError3() {
	bookingID := 1
	var booking models.Booking
	booking.AcceptedByDriver = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("MarkBookingAsCompleted", bookingID).Return(errors.New("ERROR"))

	err := s.service.MarkRideAsCompleted("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestCancelBooking() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("CancelBookingUser", bookingID).Return(nil)

	err := s.service.CancelBooking("1")
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestCancelBookingError() {
	err := s.service.CancelBooking("")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestCancelBookingError2() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1
	booking.AcceptedByDriver = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)

	err := s.service.CancelBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestCancelBookingError3() {
	bookingID := 1
	var booking models.Booking
	booking.IDCab = 1
	booking.AcceptedByDriver = true

	mockRepository.On("GetBooking", bookingID).Return(&booking, nil)
	mockRepository.On("CancelBookingUser", bookingID).Return(errors.New("ERROR"))

	err := s.service.CancelBooking("1")
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestGetBookingsByUser() {
	userID := 1
	var booking models.Booking
	bookings := []models.Booking{booking}

	mockRepository.On("GetBookingsByUserID", userID).Return(bookings, nil)

	_, err := s.service.GetBookingsByUser("1")
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestGetBookingsByUserError() {
	_, err := s.service.GetBookingsByUser("")
	assert.Error(s.T(), err)
}
