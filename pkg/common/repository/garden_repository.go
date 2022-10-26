package repository

import (
	"fmt"

	"github.com/roblesoft/plants/pkg/common/models"
)

type GardenRepository struct {
	GormRepository
}

func (r *GardenRepository) List(args map[string]any) (any, error) {
	var wc models.GardenCollection
	order := "created_at"
	err := r.db.Limit(args["limit"].(int)).Order(order).Where(fmt.Sprintf("%v > ?", order), args["after"]).Limit(args["limit"].(int)).Find(&wc).Error

	return wc, err
}

func (r *GardenRepository) Get(id any) (any, error) {
	var w *models.Garden

	err := r.db.Where("id = ?", id).First(&w).Error

	return w, err
}

func (r *GardenRepository) Create(entity any) (any, error) {
	w := entity.(*models.Garden)

	err := r.db.Create(w).Error

	return w, err
}

func (r *GardenRepository) Update(args any) (any, error) {
	w := args.(map[string]any)["entity"].(*models.Plant)

	if err := r.db.Model(w).Where("id = ?", args.(map[string]any)["id"]).Updates(w).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *GardenRepository) Delete(id any) (bool, error) {
	if err := r.db.Delete(&models.Garden{}, "id = ?", id).Error; err != nil {
		return false, err
	}

	return true, nil
}
