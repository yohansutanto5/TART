package model

type Sonar struct {
	Test
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
