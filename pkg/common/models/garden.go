package models

import (
	"time"

	"gorm.io/gorm"
)

type Garden struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index,->" json:"-"`
}

type GardenCollection []*Garden
