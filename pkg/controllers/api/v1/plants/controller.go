package plants

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/models"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/lib"
)

type PlantParams struct {
	ID         string `uri:"id" validate:"required"`
	GardenId   uint   `json:"garden_id" validate:"required"`
	CommonName string `json:"common_name"`
	Family     string `json:"family"`
	PlantClass string `json:"plant_class"`
}

type GardenParams struct {
	GardenId string `uri:"GardenId" validate:"required"`
	Name     string `json:"name"`
}

type query struct {
	After time.Time `form:"after"`
	Limit int       `form:"limit,default=10" binding:"gte=1,lte=100"`
}

func GetPlantRepository(ctx *gin.Context) repository.Repository {
	return ctx.MustGet("RepositoryRegistry").(*repository.RepositoryRegistry).MustRepository("PlantRepository")
}

func GetPlants(ctx *gin.Context) {
	p := GardenParams{}
	q := query{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	pc, err := GetPlantRepository(ctx).List(map[string]any{"after": q.After, "limit": q.Limit, "gardenId": p.GardenId, "ctx": ctx})
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &pc)
}

func GetPlant(ctx *gin.Context) {
	gp := GardenParams{}
	pp := PlantParams{}

	ctx.ShouldBindUri(&gp)

	if err := lib.Validate.Struct(gp); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.ShouldBindUri(&pp)

	if err := lib.Validate.Struct(pp); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	plant, err := GetPlantRepository(ctx).Get(map[string]any{"plantId": pp.ID, "gardenId": gp.GardenId, "ctx": ctx})
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &plant)
}

func CreatePlant(ctx *gin.Context) {
	body := models.Plant{}

	if err := ctx.BindJSON(&body); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	entity, err := GetPlantRepository(ctx).Create(&body)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	plant := entity.(*models.Plant)
	ctx.JSON(http.StatusCreated, &plant)
}

func UpdatePlant(ctx *gin.Context) {
	p := PlantParams{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	body := models.Plant{}

	if err := ctx.BindJSON(&body); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	repository := GetPlantRepository(ctx)

	_, err := repository.Update(p.ID, &body)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	plant, err := repository.Get(p.ID)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &plant)
}

func DeletePlant(ctx *gin.Context) {
	p := PlantParams{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	_, err := GetPlantRepository(ctx).Delete(p.ID)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	lib.WriteNoContent(ctx)
}
