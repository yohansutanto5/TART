package model

import "time"

type Sonar struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Artefact string    `gorm:"type:VARCHAR(20);not null;"`
	Created  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input``

type AddSonarIn struct {
	Artefact string `json:"artefact" binding:"required"`
}

type ListSonarOut struct {
	Artefact string `json:"artefact"`
	ID       int
}

func (m *Sonar) PopulateFromDTOInput(input AddSonarIn) {
	m.Artefact = input.Artefact
}
