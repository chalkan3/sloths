package sloth

import (
	"time"

	"gorm.io/gorm"
)

type SlothModel struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Family    string
	CreateAT  time.Time
	UpdatedAT time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewSlothModel(name, family string) *SlothModel {
	return &SlothModel{
		Name:   name,
		Family: family,
	}
}
