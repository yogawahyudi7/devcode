package repository

import (
	"devcode/model"
	"fmt"

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

func (td *ActivityRepository) GetAll() (data []model.Activity, err error) {

	err = td.db.Debug().Find(&data).Error

	if err != nil {
		return data, err
	}

	fmt.Println(data)
	fmt.Println(err)
	return data, err
}

func (td *ActivityRepository) GetOne(id int) (data model.Activity, err error) {

	err = td.db.Debug().Where("activity_id = ?", id).Find(&data).Error

	if err != nil {
		return data, err
	}

	fmt.Println(data)
	fmt.Println(err)
	return data, err
}

func (td *ActivityRepository) Delete(id int) (err error) {

	data := model.Todo{}

	err = td.db.Debug().Where("activity_id = ?", id).Delete(&data).Error

	if err != nil {
		return err
	}

	fmt.Println(data)
	fmt.Println(err)
	return err
}

func (td *ActivityRepository) Create(params model.Activity) (data model.Activity, err error) {

	data = model.Activity{
		Title: params.Title,
		Email: params.Email,
	}
	err = td.db.Debug().Create(&data).Error

	if err != nil {
		return data, err
	}

	fmt.Println(data)
	fmt.Println(err)
	return data, err
}

func (td *ActivityRepository) Update(id int, params model.Activity) (data model.Activity, err error) {

	err = td.db.Debug().Where("activity_id = ?", id).Find(&data).Error

	if err != nil {
		return data, err
	}

	data = model.Activity{
		Title: params.Title,
		Email: params.Email,
	}
	err = td.db.Debug().Save(&data).Error

	if err != nil {
		return data, err
	}

	fmt.Println(data)
	fmt.Println(err)
	return data, err
}
