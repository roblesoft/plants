package gardens

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/models"
	"github.com/roblesoft/plants/pkg/common/repository"
	"github.com/roblesoft/plants/pkg/lib"
)

type GardenParams struct {
	GardenID string `uri:"id" validate:"required"`
	Name     string `json:"name"`
}

type query struct {
	After time.Time `form:"after"`
	Limit int       `form:"limit,default=10" binding:"gte=1,lte=100"`
}

func GetGardenRepository(ctx *gin.Context) repository.Repository {
	return ctx.MustGet("RepositoryRegistry").(*repository.RepositoryRegistry).MustRepository("GardenRepository")
}

func GetGardens(ctx *gin.Context) {
	var q = query{}

	if err := ctx.ShouldBindQuery(&q); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	gardens, err := GetGardenRepository(ctx).List(map[string]any{"after": q.After, "limit": q.Limit})
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &gardens)
}

func GetGarden(ctx *gin.Context) {
	p := GardenParams{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	garden, err := GetGardenRepository(ctx).Get(p.GardenID)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &garden)
}

func CreateGarden(ctx *gin.Context) {
	body := models.Garden{}

	if err := ctx.BindJSON(&body); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	entity, err := GetGardenRepository(ctx).Create(&body)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	garden := entity.(*models.Garden)
	ctx.JSON(http.StatusCreated, &garden)
}

func UpdateGarden(ctx *gin.Context) {
	p := GardenParams{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	body := models.Garden{}

	if err := ctx.BindJSON(&body); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	repository := GetGardenRepository(ctx)

	_, err := repository.Update(map[string]any{"entity": &body, "id": p.GardenID})
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	garden, err := repository.Get(p.GardenID)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, &garden)
}

func DeleteGarden(ctx *gin.Context) {
	p := GardenParams{}

	ctx.ShouldBindUri(&p)

	if err := lib.Validate.Struct(p); err != nil {
		lib.HandleError(err, ctx)
		return
	}

	_, err := GetGardenRepository(ctx).Delete(p.GardenID)
	if err != nil {
		lib.HandleError(err, ctx)
		return
	}

	lib.WriteNoContent(ctx)
}
