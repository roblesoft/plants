package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers/api/v1/gardens"
	"github.com/roblesoft/plants/pkg/controllers/api/v1/plants"
	"github.com/roblesoft/plants/pkg/lib"
	"github.com/spf13/viper"
)

type Server struct {
	router *gin.Engine
}

func InitServer() *Server {
	server := &Server{}
	server.router = gin.Default()
	return server
}

func (s *Server) SetRepositoryRegistry(rr *repository.RepositoryRegistry) {
	s.router.Use(func(c *gin.Context) {
		c.Set("RepositoryRegistry", rr)
		c.Next()
	})
}

func (s *Server) getRouter() *gin.Engine {
	return s.router
}

func (s *Server) registerRoutes() {
	var router = s.router

	router.NoRoute(lib.NoRoute)

	gardensRoutes := router.Group("/gardens")
	{
		gardensRoutes.GET("", gardens.GetGardens)
		gardensRoutes.POST("", gardens.CreateGarden)
		garden := gardensRoutes.Group(":GardenId")

		garden.GET("/", gardens.GetGarden)
		garden.PATCH("/", gardens.UpdateGarden)
		garden.DELETE("/", gardens.DeleteGarden)
		plantsRoutes := garden.Group("plants")
		{
			plantsRoutes.GET("", plants.GetPlants)
			plantsRoutes.POST("", plants.CreatePlant)
			plantsRoutes.GET("/:id", plants.GetPlant)
			plantsRoutes.PATCH("/:id", plants.UpdatePlant)
			plantsRoutes.DELETE("/:id", plants.DeletePlant)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
	})
}

func SetUpRouter() *Server {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)
	db := db.Init(dbUrl)
	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})
	server := SetUpRouter()
	server.SetRepositoryRegistry(registry)
	server.registerRoutes()

	return server
}
