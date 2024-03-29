package model

import "time"

type Ticket struct {
	ID      int       `gorm:"primaryKey;autoIncrement"`
	Appcode string    `gorm:"type:VARCHAR(20);not null;"`
	Status  string    `gorm:"type:VARCHAR(20);not null;"`
	CICD    bool      `gorm:"type:bool;not null;default:true"`
	Created time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated time.Time `gorm:"type:date;default:(CURRENT_DATE)"`
	// List of testings
	UATs         []UAT         `gorm:"foreignKey:TicketID;OnDelete:CASCADE;"`
	Regressions  []Regression  `gorm:"foreignKey:TicketID;OnDelete:CASCADE;"`
	Performances []Performance `gorm:"foreignKey:TicketID;OnDelete:CASCADE;"`
	UnitTest     UnitTest      `gorm:"foreignKey:UniTestID;OnDelete:CASCADE;"`
	UniTestID    int
	Sonar        Sonar `gorm:"foreignKey:SonarID;OnDelete:CASCADE;"`
	SonarID      int
	Chaos        Chaos `gorm:"foreignKey:ChaosID;OnDelete:CASCADE;"`
	ChaosID      int
	// PAT Test
}
type AddTicketIn struct {
	Appcode      string             `json:"appcode" binding:"required"`
	CICD         bool               `json:"cicd" binding:"required"`
	UATs         []AddUATIn         `json:"uat"`
	Regressions  []AddRegressionIn  `json:"regression"`
	Performances []AddPerformanceIn `json:"performance"`
	Sonar        AddSonarIn         `json:"sonar"`
	Chaos        AddChaosIn         `json:"chaos"`
	UnitTest     AddUnitTestIn      `json:"unittest"`
}

type GetTicketOut struct {
	Appcode      string               `json:"appcode"`
	Status       string               `json:"status"`
	CICD         bool                 `json:"cicd"`
	UATs         []ListUATOut         `json:"uat"`
	Regressions  []ListRegressionOut  `json:"regression"`
	Performances []ListPerformanceOut `json:"performance"`
	Sonar        Sonar                `json:"sonar"`
	Chaos        Chaos                `json:"chaos"`
	UnitTest     UnitTest             `json:"unittest"`
}

func (m *Ticket) PopulateFromDTOInput(input AddTicketIn) {
	m.Appcode = input.Appcode
	m.Sonar.PopulateFromDTOInput(input.Sonar)
	m.UnitTest.PopulateFromDTOInput(input.UnitTest)
	m.Chaos.PopulateFromDTOInput(input.Chaos)
	
	for _, i := range input.UATs {
		m.UATs = append(m.UATs, UAT{Artefact: i.Artefact})
	}
	for _, i := range input.Regressions {
		m.Regressions = append(m.Regressions, Regression{Artefact: i.Artefact})
	}
	for _, i := range input.Performances {
		m.Performances = append(m.Performances, Performance{Artefact: i.Artefact})
	}
	m.Status = "DRAFT"
}
func (m *Ticket) ConstructGetTicketOut() (res GetTicketOut) {
	res.Appcode = m.Appcode
	res.Status = m.Status
	res.CICD = m.CICD
	res.Sonar = Sonar{Test: Test{Artefact: m.Chaos.Artefact}}
	res.UnitTest = UnitTest{Artefact: m.UnitTest.Artefact}
	res.Chaos = Chaos{Test: Test{Artefact: m.Chaos.Artefact}}

	for _, i := range m.Regressions {
		res.Regressions = append(res.Regressions, ListRegressionOut{Artefact: i.Artefact})
	}
	for _, i := range m.Performances {
		res.Performances = append(res.Performances, ListPerformanceOut{Artefact: i.Artefact})
	}
	for _, i := range m.UATs {
		res.UATs = append(res.UATs, ListUATOut{ID: i.ID, Artefact: i.Artefact})
	}

	return
}
