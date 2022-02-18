package cab

import (
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

func (m *MockRepository) GetCabsByLatLong(latitude, longitude float64) ([]models.Cab, error) {
	args := m.Called(latitude, longitude)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Cab), nil
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

func (s *ServiceSuite) TestGetNearbyCabs() {
	latitude, longitude := "10.11", "10.12"
	lat, long := 10.11, 10.12
	var cab models.Cab
	cabs := []models.Cab{cab}

	mockRepository.On("GetCabsByLatLong", lat, long).Return(cabs, nil)

	_, err := s.service.GetNearbyCabs(latitude, longitude)
	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestGetNearbyCabsError() {
	latitude, longitude := "10.1w", "10.12"

	_, err := s.service.GetNearbyCabs(latitude, longitude)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestGetNearbyCabsError2() {
	latitude, longitude := "10.11", "10.1e"

	_, err := s.service.GetNearbyCabs(latitude, longitude)
	assert.Error(s.T(), err)
}
