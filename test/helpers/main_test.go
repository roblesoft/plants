package helpers

// import "github.com/gin-gonic/gin"

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

import (
	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers"
	"github.com/spf13/viper"
)

func SetUpRouter() *controllers.Server {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)

	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})

	server := controllers.InitServer()
	server.SetRepositoryRegistry(registry)
	return server
}
