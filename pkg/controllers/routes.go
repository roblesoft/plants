package controllers

import (
	"github.com/gin-contrib/gzip"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers/plants"
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
		plantsRoutes.GET("/:uuid", plants.GetPlant)
		plantsRoutes.PATCH("/:uuid", plants.UpdatePlant)
		plantsRoutes.DELETE("/:uuid", plants.DeletePlant)
	}

	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
	})
}

func InitServer() *Server {
	return &Server{router: gin.Default()}
}

func (s *Server) Run() (err error) {
	s.registerRoutes()

	s.router.Use(gzip.Gzip(gzip.DefaultCompression))

	err = s.router.Run()
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
