package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/roblesoft/plants/pkg/common/models"
)

type PlantRepository struct {
	GormRepository
}

func GetGardenRepository(ctx *gin.Context) Repository {
	return ctx.MustGet("RepositoryRegistry").(*RepositoryRegistry).MustRepository("GardenRepository")
}

func (r *PlantRepository) List(args map[string]any) (any, error) {
	var pc models.PlantCollection

	garden, err := GetGardenRepository(args["ctx"].(*gin.Context)).Get(args["gardenId"])

	if err != nil {
		return garden, err
	}

	order := "created_at"

	err = r.db.Limit(args["limit"].(int)).Order(order).Where("garden_id = ?", &garden.(*models.Garden).ID).Limit(args["limit"].(int)).Find(&pc).Error

	return pc, err
}

func (r *PlantRepository) Get(args any) (any, error) {
	var plant *models.Plant

	garden, err := GetGardenRepository(args.(map[string]any)["ctx"].(*gin.Context)).Get(args.(map[string]any)["gardenId"])

	if err != nil {
		return garden, err
	}

	err = r.db.Where("garden_id = ? AND id = ?", &garden.(*models.Garden).ID, args.(map[string]any)["plantId"]).First(&plant).Error

	return plant, err
}

func (r *PlantRepository) Create(args any) (any, error) {
	plant := args.(map[string]any)["entity"].(*models.Plant)

	garden, err := GetGardenRepository(args.(map[string]any)["ctx"].(*gin.Context)).Get(args.(map[string]any)["gardenId"])

	if err != nil {
		return garden, err
	}

	plant.GardenId = garden.(*models.Garden).ID
	err = r.db.Create(plant).Error

	return plant, err
}

func (r *PlantRepository) Update(args any) (any, error) {
	body := args.(map[string]any)["entity"].(*models.Plant)
	garden, err := GetGardenRepository(args.(map[string]any)["ctx"].(*gin.Context)).Get(args.(map[string]any)["gardenId"])

	if err != nil {
		return false, err
	}

	if err := r.db.Model(&body).Where("garden_id = ? AND id = ?", &garden.(*models.Garden).ID, args.(map[string]any)["plantId"]).Updates(&body).Error; err != nil {
		return false, err
	}

	plant, err := r.Get(map[string]any{"plantId": args.(map[string]any)["plantId"], "gardenId": &garden.(*models.Garden).ID, "ctx": args.(map[string]any)["ctx"].(*gin.Context)})

	if err != nil {
		return false, err
	}

	return plant, nil
}

func (r *PlantRepository) Delete(args any) (bool, error) {
	garden, err := GetGardenRepository(args.(map[string]any)["ctx"].(*gin.Context)).Get(args.(map[string]any)["gardenId"])

	if err != nil {
		return false, err
	}

	if err = r.db.Delete(&models.Plant{}, "garden_id = ? AND id = ?", &garden.(*models.Garden).ID, args.(map[string]any)["plantId"]).Error; err != nil {
		return false, err
	}

	return true, nil
}
