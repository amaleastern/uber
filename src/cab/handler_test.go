package cab

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"uber/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var mockService *MockService

type MockService struct {
	mock.Mock
}

func (m *MockService) GetNearbyCabs(latitude, longitude string) ([]models.Cab, error) {
	args := m.Called(latitude, longitude)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Cab), nil
}

type HandlerSuite struct {
	suite.Suite

	handler Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	mockService = new(MockService)
	s.handler = NewHandler(mockService)
}

func (s *HandlerSuite) TestGetNearbyCabs() {
	app := fiber.New()
	app.Get("/cabs", s.handler.GetNearbyCabs)

	tests := []struct {
		description            string
		latitude               string
		longitude              string
		mockGetNearbyCabs      []models.Cab
		mockGetNearbyCabsError error
		expectedCode           int
		expectedBody           string
	}{
		{
			description:            "GetNearbyCabs success",
			latitude:               "10.11",
			longitude:              "10.12",
			mockGetNearbyCabs:      []models.Cab{{}},
			mockGetNearbyCabsError: nil,
			expectedCode:           http.StatusOK,
			expectedBody:           `[{"id":0,"NoOfSeats":0,"DriverName":"","CarModel":"","IDLocation":0,"Active":false}]`,
		},
		{
			description:            "GetNearbyCabs error",
			latitude:               "10.11",
			longitude:              "10.12",
			mockGetNearbyCabs:      nil,
			mockGetNearbyCabsError: errors.New("ERROR"),
			expectedCode:           http.StatusBadRequest,
			expectedBody:           `{"error":"ERROR"}`,
		},
		{
			description:  "GetNearbyCabs validation error",
			latitude:     "10.11",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"latitude and longitude are required"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/cabs?latitude="+test.latitude+"&longitude="+test.longitude, nil)

		mockService.On("GetNearbyCabs", test.latitude, test.longitude).Return(test.mockGetNearbyCabs, test.mockGetNearbyCabsError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}
