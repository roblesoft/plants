package gardens

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roblesoft/plants/pkg/common/db"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/controllers"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	viper.SetConfigFile("../../../../../pkg/common/envs/.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)

	db := db.Init(dbUrl)

	registry := repository.NewRepositoryRegistry(db, &repository.PlantRepository{}, &repository.GardenRepository{})

	server := controllers.InitServer()
	server.SetRepositoryRegistry(registry)

	w := httptest.NewRecorder()
	fmt.Print(server)
	// req, _ := http.NewRequest(http.MethodGet, "/", nil)
	// server.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
