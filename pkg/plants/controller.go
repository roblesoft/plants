package plants

import (
    "github.com/gin-gonic/gin"

    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/plants")
    routes.POST("/", h.AddPlant)
    routes.GET("/", h.GetPlants)
    routes.GET("/:id", h.GetPlant)
    routes.PUT("/:id", h.UpdatePlant)
    routes.DELETE("/:id", h.DeletePlant)
}
