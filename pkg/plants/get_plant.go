
package plants

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/roblesoft/plants/pkg/common/models"
)

func (h handler) GetPlant(c *gin.Context) {
    id := c.Param("id")

    var plant models.Plant

    if result := h.DB.First(&plant, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &plant)
}