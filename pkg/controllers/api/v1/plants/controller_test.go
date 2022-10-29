package plants

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
	"github.com/roblesoft/plants/pkg/controllers/api/v1/gardens"
	"github.com/roblesoft/plants/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetPlants(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	body := models.Garden{Name: "Uriel"}

	r.GET("/gardens/:GardenId/plants", GetPlants)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{})
	ctx.Set("RepositoryRegistry", registry)
	entity, _ := repository.GetGardenRepository(ctx).Create(&body)

	r.GET("/gardens", gardens.GetGardens)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/gardens/%d/plants", entity.(*models.Garden).ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPlantsGardenNotFound(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router

	r.GET("/gardens/:GardenId/plants", GetPlants)

	w := httptest.NewRecorder()

	r.GET("/gardens", gardens.GetGardens)
	req, _ := http.NewRequest("GET", "/gardens/0/plants", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPlantSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	garden_body := models.Garden{Name: "Uriel"}
	plant_body := models.Plant{CommonName: "Monstera"}

	r.GET("/gardens/:GardenId/plants/:id", GetPlant)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{}, &repository.PlantRepository{})
	ctx.Set("RepositoryRegistry", registry)

	garden, _ := repository.GetGardenRepository(ctx).Create(&garden_body)
	plant, _ := mock.GetPlantRepository(ctx).Create(map[string]any{"entity": &plant_body, "gardenId": garden.(*models.Garden).ID, "ctx": ctx})

	req, _ := http.NewRequest("GET", fmt.Sprintf("/gardens/%d/plants/%d", garden.(*models.Garden).ID, plant.(*models.Plant).ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostPlantSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	garden_body := models.Garden{Name: "Uriel"}
	plant_body := models.Plant{CommonName: "Monstera"}

	r.POST("/gardens/:GardenId/plants", CreatePlant)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{}, &repository.PlantRepository{})
	ctx.Set("RepositoryRegistry", registry)

	garden, _ := repository.GetGardenRepository(ctx).Create(&garden_body)

	jsonValue, _ := json.Marshal(plant_body)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/gardens/%d/plants", garden.(*models.Garden).ID), bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)
	plant := models.Plant{}
	if err := json.Unmarshal(w.Body.Bytes(), &plant); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	assert.Equal(t, plant.CommonName, plant_body.CommonName)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatePlantSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	garden_body := models.Garden{Name: "Uriel"}
	plant_body := models.Plant{CommonName: "Monstera"}
	updatedPlantBody := models.Plant{CommonName: "Updated Monstera"}

	r.PATCH("/gardens/:GardenId/plants/:id", UpdatePlant)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{}, &repository.PlantRepository{})
	ctx.Set("RepositoryRegistry", registry)

	garden, _ := repository.GetGardenRepository(ctx).Create(&garden_body)
	plant, _ := mock.GetPlantRepository(ctx).Create(map[string]any{"entity": &plant_body, "gardenId": garden.(*models.Garden).ID, "ctx": ctx})
	fmt.Print(plant)

	jsonValue, _ := json.Marshal(updatedPlantBody)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/gardens/%d/plants/%d", garden.(*models.Garden).ID, plant.(*models.Plant).ID), bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &plant); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	assert.Equal(t, plant.(*models.Plant).CommonName, updatedPlantBody.CommonName)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePlantSuccessfully(t *testing.T) {
	mock.SetUpRouter()
	r := mock.Router
	garden_body := models.Garden{Name: "Uriel"}
	plant_body := models.Plant{CommonName: "Monstera"}

	r.DELETE("/gardens/:GardenId/plants/:id", DeletePlant)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	registry := repository.NewRepositoryRegistry(mock.DB, &repository.GardenRepository{}, &repository.PlantRepository{})
	ctx.Set("RepositoryRegistry", registry)

	garden, _ := repository.GetGardenRepository(ctx).Create(&garden_body)
	plant, _ := mock.GetPlantRepository(ctx).Create(map[string]any{"entity": &plant_body, "gardenId": garden.(*models.Garden).ID, "ctx": ctx})

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/gardens/%d/plants/%d", garden.(*models.Garden).ID, plant.(*models.Plant).ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
