package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var articleJSON = `{"tittle":"Jon Snow","body":"jonksjhfkjahlskjdfhlakjshdkfjhaljshdfkjahsdf"}`

func TestCreateArticle(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/article", strings.NewReader(articleJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, CreateArticle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "["+articleJSON+"]", rec.Body.String())
	}
}

/* func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
} */
