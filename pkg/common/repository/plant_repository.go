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

func (r *PlantRepository) Get(id any) (any, error) {
	var w *models.Plant

	err := r.db.Where("id = ?", id).First(&w).Error

	return w, err
}

func (r *PlantRepository) Create(entity any) (any, error) {
	w := entity.(*models.Plant)

	err := r.db.Create(w).Error

	return w, err
}

func (r *PlantRepository) Update(id any, entity any) (bool, error) {
	w := entity.(*models.Plant)

	if err := r.db.Model(w).Where("id = ?", id).Updates(w).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *PlantRepository) Delete(id any) (bool, error) {
	if err := r.db.Delete(&models.Plant{}, "id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}
