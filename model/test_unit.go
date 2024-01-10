package model

import "time"

type UnitTest struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Artefact string    `gorm:"type:VARCHAR(20);not null;"`
	Created  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input``

type AddUnitTestIn struct {
	Artefact string `json:"artefact" binding:"required"`
}

type ListUnitTestOut struct {
	Artefact string `json:"artefact"`
	ID       int
}

func (m *UnitTest) PopulateFromDTOInput(input AddUnitTestIn) {
	m.Artefact = input.Artefact
}
