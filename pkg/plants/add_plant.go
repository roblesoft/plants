package plants

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/roblesoft/plants/pkg/common/models"
)

type AddPlantRequestBody struct {
	CommonName string `json:"common_name"`
	Family     string `json:"family"`
	PlantClass string `json:"plant_class"`
}

func (h handler) AddPlant(c *gin.Context) {
    body := AddPlantRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var plant models.Plant

    plant.CommonName = body.CommonName
    plant.Family = body.Family
    plant.PlantClass = body.PlantClass

    if result := h.DB.Create(&plant); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &plant)
}