package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/spf13/viper"
)

type Server struct {
	router *gin.Engine
}

var Router *gin.Engine

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

func SetUpRouter() {
	viper.SetConfigFile("../../../../common/envs/.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)
	db := db.Init(dbUrl)
	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})
	server := InitServer()
	server.SetRepositoryRegistry(registry)
	Router = server.router
}
