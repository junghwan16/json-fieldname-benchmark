package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleShortJSON(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/short", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handleShortJSON(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp ShortResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.Equal(t, "Alice", resp.A)
		assert.True(t, resp.C)
	}
}

func TestHandleLongJSON(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/long", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handleLongJSON(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp LongResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.Equal(t, "Alice Johnson", resp.CustomerFullName)
		assert.True(t, resp.IsEligibleForPromotionalAd)
	}
}
