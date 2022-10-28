package gardens

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roblesoft/plants/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetGardens(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	r.GET("/gardens", GetGardens)
	req, _ := http.NewRequest("GET", "/gardens", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
