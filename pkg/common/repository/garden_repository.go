package repository

import (
	"fmt"
	"time"

	"github.com/roblesoft/plants/pkg/common/models"
)

type GardenRepository struct {
	GormRepository
}

func (r *GardenRepository) List(after time.Time, limit int) (any, error) {
	var wc models.GardenCollection
	order := "created_at"
	err := r.db.Limit(limit).Order(order).Where(fmt.Sprintf("%v > ?", order), after).Limit(limit).Find(&wc).Error

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

func (r *GardenRepository) Update(id any, entity any) (bool, error) {
	w := entity.(*models.Garden)

	if err := r.db.Model(w).Where("id = ?", id).Updates(w).Error; err != nil {
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
