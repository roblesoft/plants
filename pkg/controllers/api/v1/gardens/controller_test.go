package gardens

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := SetUpRouter()
	r.router.GET("/gardens", GetGardens)
	req, _ := http.NewRequest("GET", "/gardens", nil)
	w := httptest.NewRecorder()
	r.getRouter().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

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

func SetUpRouter() *Server {
	viper.SetConfigFile("../../../../common/envs/.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)
	db := db.Init(dbUrl)
	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})
	server := SetUpRouter()
	server.SetRepositoryRegistry(registry)

	return server
}
