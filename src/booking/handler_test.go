package booking

import (
	"bytes"
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

func (m *MockService) GetBookingsByUser(userID string) ([]models.Booking, error) {
	args := m.Called(userID)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Booking), nil
}

func (m *MockService) CreateBooking(booking *models.Booking) error {
	args := m.Called(booking)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockService) CancelBooking(bookingID string) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockService) AcceptBooking(bookingID string) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (m *MockService) MarkRideAsCompleted(bookingID string) error {
	args := m.Called(bookingID)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
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
	app.Get("/bookings/:UserID", s.handler.GetBookingsByUser)

	tests := []struct {
		description                string
		userID                     string
		mockGetBookingsByUser      []models.Booking
		mockGetBookingsByUserError error
		expectedCode               int
		expectedBody               string
	}{
		{
			description:                "GetBookingsByUser success",
			userID:                     "1",
			mockGetBookingsByUser:      []models.Booking{{}},
			mockGetBookingsByUserError: nil,
			expectedCode:               http.StatusOK,
			expectedBody:               `[{"id":0,"id_user":0,"id_cab":0,"id_from_location":0,"id_to_location":0,"cancelled_by_user":false,"cancelled_by_driver":false,"accepted_by_driver":false,"completed":false,"booking_date":"0001-01-01T00:00:00Z"}]`,
		},
		{
			description:                "GetBookingsByUser error",
			userID:                     "1",
			mockGetBookingsByUser:      nil,
			mockGetBookingsByUserError: errors.New("ERROR"),
			expectedCode:               http.StatusOK,
			expectedBody:               `{"error":"ERROR"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/bookings/"+test.userID, nil)

		mockService.On("GetBookingsByUser", test.userID).Return(test.mockGetBookingsByUser, test.mockGetBookingsByUserError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}

func (s *HandlerSuite) TestBookRide() {
	app := fiber.New()
	app.Post("/book", s.handler.BookRide)

	tests := []struct {
		description            string
		requestBody            string
		booking                models.Booking
		mockCreateBookingError error
		expectedCode           int
		expectedBody           string
	}{
		{
			description: "BookRide success",
			requestBody: `{"id_cab": 1,"id_user": 1,"id_from_location": 1,"id_to_location": 2}`,
			booking: models.Booking{
				IDCab:          1,
				IDUser:         1,
				IDFromLocation: 1,
				IDToLocation:   2,
			},
			mockCreateBookingError: nil,
			expectedCode:           http.StatusOK,
			expectedBody:           `{"Status":"Booked"}`,
		},
		{
			description: "BookRide error",
			requestBody: `{"id_cab": 1,"id_user": 1,"id_from_location": 1,"id_to_location": 2}`,
			booking: models.Booking{
				IDCab:          1,
				IDUser:         1,
				IDFromLocation: 1,
				IDToLocation:   2,
			},
			mockCreateBookingError: errors.New("ERROR"),
			expectedCode:           http.StatusBadRequest,
			expectedBody:           `{"error":"ERROR"}`,
		},
		{
			description:  "BookRide validation error",
			requestBody:  `{"id_cab": 1,"id_user": 1,"id_from_location": 1}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"id_to_location: non zero value required"}`,
		},
		{
			description:  "BookRide BodyParser error",
			requestBody:  `{`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"json: unexpected end of JSON input: {"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/book", bytes.NewBuffer([]byte(test.requestBody)))
		req.Header.Add("Content-Type", "application/json")
		mockService.On("CreateBooking", &test.booking).Return(test.mockCreateBookingError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}

func (s *HandlerSuite) TestCancelBooking() {
	app := fiber.New()
	app.Put("/booking/:BookingID/cancel", s.handler.CancelBooking)

	tests := []struct {
		description            string
		bookingID              string
		mockCancelBookingError error
		expectedCode           int
		expectedBody           string
	}{
		{
			description:            "CancelBooking success",
			bookingID:              "1",
			mockCancelBookingError: nil,
			expectedCode:           http.StatusOK,
			expectedBody:           `{"Status":"Updated"}`,
		},
		{
			description:            "CancelBooking error",
			bookingID:              "1",
			mockCancelBookingError: errors.New("ERROR"),
			expectedCode:           http.StatusBadRequest,
			expectedBody:           `{"error":"ERROR"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPut, "/booking/"+test.bookingID+"/cancel", nil)

		mockService.On("CancelBooking", test.bookingID).Return(test.mockCancelBookingError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}

func (s *HandlerSuite) TestAcceptBooking() {
	app := fiber.New()
	app.Put("/booking/:BookingID/accept", s.handler.AcceptBooking)

	tests := []struct {
		description            string
		bookingID              string
		mockAcceptBookingError error
		expectedCode           int
		expectedBody           string
	}{
		{
			description:            "AcceptBooking success",
			bookingID:              "1",
			mockAcceptBookingError: nil,
			expectedCode:           http.StatusOK,
			expectedBody:           `{"Status":"Updated"}`,
		},
		{
			description:            "AcceptBooking error",
			bookingID:              "1",
			mockAcceptBookingError: errors.New("ERROR"),
			expectedCode:           http.StatusBadRequest,
			expectedBody:           `{"error":"ERROR"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPut, "/booking/"+test.bookingID+"/accept", nil)

		mockService.On("AcceptBooking", test.bookingID).Return(test.mockAcceptBookingError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}

func (s *HandlerSuite) TestMarkRideAsCompleted() {
	app := fiber.New()
	app.Put("/booking/:BookingID/complete", s.handler.MarkRideAsCompleted)

	tests := []struct {
		description                  string
		bookingID                    string
		mockMarkRideAsCompletedError error
		expectedCode                 int
		expectedBody                 string
	}{
		{
			description:                  "MarkRideAsCompleted success",
			bookingID:                    "1",
			mockMarkRideAsCompletedError: nil,
			expectedCode:                 http.StatusOK,
			expectedBody:                 `{"Status":"Updated"}`,
		},
		{
			description:                  "AcceptBooking error",
			bookingID:                    "1",
			mockMarkRideAsCompletedError: errors.New("ERROR"),
			expectedCode:                 http.StatusBadRequest,
			expectedBody:                 `{"error":"ERROR"}`,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPut, "/booking/"+test.bookingID+"/complete", nil)

		mockService.On("MarkRideAsCompleted", test.bookingID).Return(test.mockMarkRideAsCompletedError).Once()

		resp, err := app.Test(req, -1)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedCode, resp.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nilf(s.T(), err, test.description)

		assert.Equalf(s.T(), test.expectedBody, string(body), test.description)
	}
}
