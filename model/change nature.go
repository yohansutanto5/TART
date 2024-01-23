package model

type ChangeNature struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:VARCHAR(20);not null;" convert:"Profile.Name"`
}

// DTO input

type AddChangeNatureIn struct {
	Name string `json:"Name" binding:"required"`
}

func (m *ChangeNature) PopulateFromDTOInput(input AddChangeNatureIn) {
	m.Name = input.Name
}
