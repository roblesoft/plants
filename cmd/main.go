package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/plants"
	"github.com/roblesoft/plants/pkg/common/db"
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

	r.Run(port)
}
