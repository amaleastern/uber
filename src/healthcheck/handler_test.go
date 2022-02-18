package healthcheck

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	description := "heathcheck"
	req := httptest.NewRequest("GET", "/", nil)

	app := fiber.New()
	app.Get("/", Index)

	res, err := app.Test(req, -1)
	assert.Nilf(t, err, description)

	assert.Equalf(t, http.StatusOK, res.StatusCode, description)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, description)

	expectedBody := `{"Status":"OK"}`

	assert.Equalf(t, expectedBody, string(body), description)
}
