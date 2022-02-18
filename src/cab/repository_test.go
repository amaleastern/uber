package cab

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

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

func (s *RepositorySuite) TestGetCabsByLatLongError() {
	radius := .1
	latitude, longitude := 10.11, 10.12
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT `cab`.`id`,`cab`.`no_of_seats`,`cab`.`driver_name`,`cab`.`car_model`,`cab`.`id_location`,`cab`.`active` FROM `cab` JOIN location ON location.id = cab.id_location WHERE (location.latitude BETWEEN ? AND ?) AND (location.longitude BETWEEN ? AND ?) AND active = ?")).
		WithArgs(latitude-radius, latitude+radius, longitude-radius, longitude+radius, true).
		WillReturnError(errors.New("ERROR"))

	_, err := s.repository.GetCabsByLatLong(latitude, longitude)

	assert.Error(s.T(), err)
}

func (s *RepositorySuite) TestGetCabsByLatLong() {
	radius := .1
	latitude, longitude := 10.11, 10.12

	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT `cab`.`id`,`cab`.`no_of_seats`,`cab`.`driver_name`,`cab`.`car_model`,`cab`.`id_location`,`cab`.`active` FROM `cab` JOIN location ON location.id = cab.id_location WHERE (location.latitude BETWEEN ? AND ?) AND (location.longitude BETWEEN ? AND ?) AND active = ?")).
		WithArgs(latitude-radius, latitude+radius, longitude-radius, longitude+radius, true).
		WillReturnRows(sqlmock.NewRows(nil))

	_, err := s.repository.GetCabsByLatLong(latitude, longitude)

	assert.NoError(s.T(), err)
}
