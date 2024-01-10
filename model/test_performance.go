package model

import "time"

type Performance struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	Artefact string    `gorm:"type:VARCHAR(20);not null;"`
	TicketID int       `gorm:"index;not null"`
	Ticket   Ticket    `gorm:"constraint:onupdate"`
	Created  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated  time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input``

type AddPerformanceIn struct {
	Artefact string `json:"artefact" binding:"required"`
}

type ListPerformanceOut struct {
	Artefact string `json:"artefact"`
	ID       int
}

func (m *Performance) PopulateFromDTOInput(input AddPerformanceIn) {
	m.Artefact = input.Artefact
}
