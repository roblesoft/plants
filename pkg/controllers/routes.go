package controllers

import (
	"github.com/gin-contrib/gzip"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers/api/v1/gardens"
	"github.com/roblesoft/plants/pkg/controllers/api/v1/plants"
	"github.com/roblesoft/plants/pkg/lib"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) registerRoutes() {
	var router = s.router

	router.NoRoute(lib.NoRoute)

	plantsRoutes := router.Group("/plants")
	{
		plantsRoutes.GET("", plants.GetPlants)
		plantsRoutes.POST("", plants.CreatePlant)
		plantsRoutes.GET("/:id", plants.GetPlant)
		plantsRoutes.PATCH("/:id", plants.UpdatePlant)
		plantsRoutes.DELETE("/:id", plants.DeletePlant)
	}

	gardensRoutes := router.Group("/gardens")
	{
		gardensRoutes.GET("", gardens.GetGardens)
		gardensRoutes.POST("", gardens.CreateGarden)
		gardensRoutes.GET("/:id", gardens.GetGarden)
		gardensRoutes.GET("/:id/plants", gardens.GetPlants)
		gardensRoutes.PATCH("/:id", gardens.UpdateGarden)
		gardensRoutes.DELETE("/:id", gardens.DeleteGarden)
	}

	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
	})
}

func InitServer() *Server {
	return &Server{router: gin.Default()}
}

func (s *Server) Run(port string) (err error) {
	s.registerRoutes()

	s.router.Use(gzip.Gzip(gzip.DefaultCompression))

	err = s.router.Run(port)
	return
}

func (s *Server) SetRepositoryRegistry(rr *repository.RepositoryRegistry) {
	s.router.Use(func(c *gin.Context) {
		c.Set("RepositoryRegistry", rr)
		c.Next()
	})
}

func (s *Server) SetLogger(logger *zap.Logger) {
	s.router.Use(func(c *gin.Context) {
		c.Set("Logger", logger)
		c.Next()
	})

	// s.router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// s.router.Use(ginzap.RecoveryWithZap(logger, true))
}
