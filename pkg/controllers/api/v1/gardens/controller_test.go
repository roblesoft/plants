package gardens

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/models"
	"github.com/roblesoft/plants/pkg/common/repository"
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

func TestGetGardenNotFound(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	r.GET("/gardens", GetGarden)
	req, _ := http.NewRequest("GET", "/gardens/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetGardenSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	body := models.Garden{Name: "Uriel"}

	r.GET("/gardens/2", GetGarden)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{})
	ctx.Set("RepositoryRegistry", registry)
	entity, _ := GetGardenRepository(ctx).Create(&body)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/gardens/%d", entity.(*models.Garden).ID), nil)
	r.ServeHTTP(w, req)
	assert.NotNil(t, entity)
	assert.Equal(t, http.StatusOK, w.Code)
}
