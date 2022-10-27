package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roblesoft/plants/pkg/controllers/api/v1/gardens"
	"github.com/roblesoft/plants/test/helpers"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := helpers.SetUpRouter()
	router := r.getRouter()
	router.GET("/", gardens.GetGardens)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
