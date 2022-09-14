
package plants

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/roblesoft/plants/pkg/common/models"
)

func (h handler) GetPlants(c *gin.Context) {
    var plants []models.Plant

    if result := h.DB.Find(&plants); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &plants)
}
