package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// ChangeNatureService defines the interface for managing ChangeNatures.
type ChangeNatureService interface {
	Insert(ChangeNature *model.ChangeNature) *error.Error
	GetList() ([]model.ChangeNature, *error.Error)
}

type ChangeNatureServiceImpl struct {
	db *db.DataStore
}

func NewChangeNatureService(db *db.DataStore) ChangeNatureService {
	return &ChangeNatureServiceImpl{db: db}
}

// Function Implementation

func (s ChangeNatureServiceImpl) GetList() ([]model.ChangeNature, *error.Error) {
	return s.db.GetListChangeNature()
}

func (s *ChangeNatureServiceImpl) Insert(ChangeNature *model.ChangeNature) *error.Error {
	return s.db.InsertChangeNature(ChangeNature)
}
