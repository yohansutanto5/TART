package db

import (
	"app/model"
	"app/pkg/error"
)

func (d *DataStore) GetListChangeNature() (ChangeNatures []model.ChangeNature, err *error.Error) {
	e := d.Db.Find(&ChangeNatures).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertChangeNature(ChangeNature *model.ChangeNature) (err *error.Error) {
	e := d.Db.Create(ChangeNature).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) DeleteChangeNatureByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.ChangeNature{}).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdateChangeNature(ChangeNature *model.ChangeNature) (err *error.Error) {
	e := d.Db.Save(&ChangeNature).Error
	err.ParseMysqlError(e)
	return
}
