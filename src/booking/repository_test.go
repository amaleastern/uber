package booking

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"uber/src/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepositorySuite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository Repository
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupTest() {
	var db *sql.DB
	var err error

	db, s.mock, err = sqlmock.New()
	assert.NoError(s.T(), err)

	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(s.T(), err)

	s.repository = NewRepository(s.DB)
}

func (s *RepositorySuite) TestGetBookingError() {
	bookingID := 1
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `booking` WHERE id = ? ORDER BY `booking`.`id` LIMIT 1")).
		WithArgs(bookingID).
		WillReturnError(errors.New("ERROR"))

	_, err := s.repository.GetBooking(bookingID)

	assert.Error(s.T(), err)
}

func (s *RepositorySuite) TestGetBookingsByUserID() {
	userID := 1
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `booking` WHERE id_user = ?")).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err := s.repository.GetBookingsByUserID(userID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestAcceptBooking() {
	bookingID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `booking` SET `accepted_by_driver`=? WHERE `id` = ?")).
		WithArgs(true, bookingID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.AcceptBooking(bookingID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestCancelBookingUser() {
	bookingID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `booking` SET `cancelled_by_user`=? WHERE `id` = ?")).
		WithArgs(true, bookingID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.CancelBookingUser(bookingID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestCancelBookingDriver() {
	bookingID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `booking` SET `cancelled_by_driver`=? WHERE `id` = ?")).
		WithArgs(true, bookingID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.CancelBookingDriver(bookingID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestMarkCabAsBusy() {
	cabID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `cab` SET `active`=? WHERE `id` = ?")).
		WithArgs(false, cabID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.MarkCabAsBusy(cabID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestMarkCabAsActive() {
	cabID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `cab` SET `active`=? WHERE `id` = ?")).
		WithArgs(true, cabID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.MarkCabAsActive(cabID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestMarkBookingAsCompleted() {
	bookingID := 1
	s.mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `booking` SET `completed`=? WHERE `id` = ?")).
		WithArgs(true, bookingID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.MarkBookingAsCompleted(bookingID)

	assert.NoError(s.T(), err)
}

func (s *RepositorySuite) TestCreateBooking() {
	var booking models.Booking
	s.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `booking` (`id_user`,`id_cab`,`id_from_location`,`id_to_location`,`cancelled_by_user`,`cancelled_by_driver`,`accepted_by_driver`,`completed`,`booking_date`) VALUES (?,?,?,?,?,?,?,?,?)")).
		WithArgs(booking.IDUser, booking.IDCab, booking.IDFromLocation, booking.IDToLocation, booking.CancelledByUser, booking.CancelledByDriver, booking.AcceptedByDriver, booking.Completed, booking.BookingDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.CreateBooking(&booking)

	assert.NoError(s.T(), err)
}
