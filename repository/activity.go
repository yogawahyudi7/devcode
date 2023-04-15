package repository

import (
	"devcode/model"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{
		db: db,
	}
}

func (db *ActivityRepository) GetAll() (data []model.Activity, err error) {

	err = db.db.Debug().Find(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (db *ActivityRepository) GetOne(id int) (data model.Activity, err error) {

	err = db.db.Debug().Where("activity_id = ?", id).Find(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (db *ActivityRepository) Delete(id int) (rowAffected int64, err error) {

	data := model.Activity{}

	query := db.db.Debug().Where("activity_id = ?", id).Delete(&data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, query.Error
}

func (db *ActivityRepository) Create(params model.Activity) (data model.Activity, err error) {

	data = model.Activity{
		Title: params.Title,
		Email: params.Email,
	}
	err = db.db.Debug().Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (db *ActivityRepository) Update(id int, params model.Activity) (rowAffected int64, err error) {

	data := map[string]interface{}{
		"title": params.Title,
		"email": params.Email,
	}

	query := db.db.Debug().Model(&params)

	query = query.Where("activity_id = ?", id)

	query = query.Updates(data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, err
}
