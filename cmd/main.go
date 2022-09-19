package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers/plants"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		fmt.Printf("%d variable", 500)
		c.Writer.WriteHeader(200)
	})

	plants.RegisterRoutes(r, h)
	registry := repository.NewRepositoryRegistry(
		h,
		&repository.PlantRepository{},
	)

	plants.SetRepositoryRegistry(r, registry)

	r.Run(port)
}
