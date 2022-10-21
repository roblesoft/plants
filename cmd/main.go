package main

import (
	"go.uber.org/zap"

	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	logger, _ := zap.NewDevelopment()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)

	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})

	server := controllers.InitServer()
	server.SetRepositoryRegistry(registry)
	server.SetLogger(logger)
	server.Run(port)
}
