package db

import (
	"app/model"
)

type TicketRepository struct {
}

func (d *DataStore) GetListTicket() (Tickets []model.Ticket, err error) {
	err = d.Db.Preload("UATs").Find(&Tickets).Error
	return
}

func (d *DataStore) InsertTicket(Ticket *model.Ticket) (err error) {
	err = d.Db.Create(Ticket).Error
	return
}

func (d *DataStore) DeleteTicketByID(id int) (err error) {
	err = d.Db.Where("id = ?", id).Delete(&model.Ticket{}).Error
	return
}

func (d *DataStore) UpdateTicket(Ticket *model.Ticket) (err error) {
	err = d.Db.Save(&Ticket).Error
	return
}
