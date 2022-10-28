package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Plant struct {
	ID         uint           `json:"id"`
	GardenId   uint           `json:"garden_id"`
	CommonName string         `json:"common_name"`
	Family     string         `json:"family"`
	PlantClass string         `json:"plant_class"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index,->" json:"-"`
}

type PlantCollection []*Plant

func (p *Plant) FullInformation() string {
	return fmt.Sprintf("Name: %s\n, Family: %s\n, Plant class: %s",
		p.CommonName, p.Family, p.PlantClass)
}
