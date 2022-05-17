package jump

import (
	"time"

	"gorm.io/gorm"
)

type Jump struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Jumped    bool
	JumpedAT  time.Time
	CreateAT  time.Time
	UpdatedAT time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewJump(name string, jumped bool, jumpedAT time.Time) *Jump {
	return &Jump{
		Name:     name,
		Jumped:   jumped,
		JumpedAT: jumpedAT,
	}
}
