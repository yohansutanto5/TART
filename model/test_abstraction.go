package model

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Artefact string    `gorm:"type:VARCHAR(20);not null;"`
	Created  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
}

func (m *Test) PopulateFromDTOInput(input AddChaosIn) {
	m.Artefact = input.Artefact
}
