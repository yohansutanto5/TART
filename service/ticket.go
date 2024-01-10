package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// TicketService defines the interface for managing Tickets.
type TicketService interface {
	Insert(Ticket *model.Ticket) *error.Error
	GetList() ([]model.Ticket, *error.Error)
}

type TicketServiceImpl struct {
	db *db.DataStore
}

func NewTicketService(db *db.DataStore) TicketService {
	return &TicketServiceImpl{db: db}
}

func (s TicketServiceImpl) GetList() ([]model.Ticket, *error.Error) {
	err := &error.Error{}
	result, e := s.db.GetListTicket()
	if e != nil {
		err.ParseMysqlError(e)
		return nil, err
	} else {
		return result, nil
	}

}

func (s *TicketServiceImpl) Insert(Ticket *model.Ticket) *error.Error {
	err := &error.Error{}
	e := s.db.InsertTicket(Ticket)
	if e != nil {
		err.ParseMysqlError(e)
		return err
	}
	return nil
}
