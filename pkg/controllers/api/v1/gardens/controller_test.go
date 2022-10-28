package gardens

import (
	"bytes"
	"encoding/json"
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

	r.GET("/gardens/:GardenId", GetGarden)

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

func TestUpdateGardenSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{})
	r := mock.Router
	garden := models.Garden{Name: "Uriel"}
	body := models.Garden{Name: "Uriel Update"}

	r.PATCH("/gardens/:GardenId", UpdateGarden)

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("RepositoryRegistry", registry)
	createdGarden, _ := GetGardenRepository(ctx).Create(&garden)

	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/gardens/%d", createdGarden.(*models.Garden).ID), bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)
	if err := json.Unmarshal(w.Body.Bytes(), &garden); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	assert.Equal(t, garden.Name, body.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteGardenSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	body := models.Garden{Name: "Uriel"}

	r.DELETE("/gardens/:GardenId", DeleteGarden)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{})
	ctx.Set("RepositoryRegistry", registry)
	entity, _ := GetGardenRepository(ctx).Create(&body)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/gardens/%d", entity.(*models.Garden).ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
