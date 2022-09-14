package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model        // adds ID, created_at etc.
	CommonName string `json:"common_name"`
	Family     string `json:"family"`
	PlantClass string `json:"plant_class"`
}

func (p *Plant) FullInformation() string {
	return fmt.Sprintf("Name: %s\n, Family: %s\n, Plant class: %s",
	                   p.CommonName, p.Family, p.PlantClass)
}
