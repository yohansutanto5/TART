package model

type Chaos struct {
	Test
}

type AddChaosIn struct {
	Artefact string `json:"artefact" binding:"required"`
}

type ListChaosOut struct {
	Artefact string `json:"artefact"`
	ID       int
}

func (m *Chaos) PopulateFromDTOInput(input AddChaosIn) {
	m.Artefact = input.Artefact
}
