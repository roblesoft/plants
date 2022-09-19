package repository

import (
	"fmt"
	"time"

	"github.com/roblesoft/plants/pkg/common/models"
)

type PlantRepository struct {
	GormRepository
}

func (r *PlantRepository) List(after time.Time, limit int) (any, error) {
	var wc models.PlantCollection
	order := "created_at"
	err := r.db.Limit(limit).Order(order).Where(fmt.Sprintf("%v > ?", order), after).Limit(limit).Find(&wc).Error

	return wc, err
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
