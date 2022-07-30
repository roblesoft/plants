package models

import "gorm.io/gorm"

type Plant struct {
	gorm.Model        // adds ID, created_at etc.
	CommonName string `json:"common_name"`
	Family     string `json:"family"`
	PlantClass string `json:"plant_class"`
}
