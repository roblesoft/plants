package plants

import (
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/lib"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// h := &handler{
	// 	DB: db,
	// }

	plants := r.Group("/plants")
	r.NoRoute(lib.NoRoute)

	{
		plants.GET("", GetPlants)
		plants.POST("", CreatePlant)
		plants.GET("/:uuid", GetPlant)
		plants.PATCH("/:uuid", UpdatePlant)
		plants.DELETE("/:uuid", DeletePlant)
	}
}

func SetRepositoryRegistry(r *gin.Engine, rr *repository.RepositoryRegistry) {
	r.Use(func(c *gin.Context) {
		c.Set("RepositoryRegistry", rr)
		c.Next()
	})
}

// func (s *Server) SetLogger(logger *zap.Logger) {
// 	s.router.Use(func(c *gin.Context) {
// 		c.Set("Logger", logger)
// 		c.Next()
// 	})

// s.router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
// s.router.Use(ginzap.RecoveryWithZap(logger, true))
// }
