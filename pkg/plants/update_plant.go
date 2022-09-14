
package plants

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/roblesoft/plants/pkg/common/models"
)

type UpdatePlantRequestBody struct {
	CommonName string `json:"common_name"`
	Family     string `json:"family"`
	PlantClass string `json:"plant_class"`
}

func (h handler) UpdatePlant(c *gin.Context) {
    id := c.Param("id")
    body := UpdatePlantRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var plant models.Plant

    if result := h.DB.First(&plant, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    plant.CommonName = body.CommonName
    plant.Family = body.Family
    plant.PlantClass = body.PlantClass

    h.DB.Save(&plant)

    c.JSON(http.StatusOK, &plant)
}
