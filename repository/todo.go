package repository

import (
	"devcode/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (db *TodoRepository) GetAll(id interface{}) (data []model.Todo, err error) {

	if id != nil {
		err = db.db.Debug().Where("activity_group_id = ?", id).Find(&data).Error
	} else {
		err = db.db.Debug().Find(&data).Error
	}

	if err != nil {
		return data, err
	}

	return data, err
}

func (db *TodoRepository) GetOne(id int) (data model.Todo, err error) {

	err = db.db.Debug().Where("todo_id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (db *TodoRepository) Delete(id int) (rowAffected int64, err error) {

	data := model.Todo{}

	query := db.db.Debug().Where("todo_id = ?", id).Delete(&data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, query.Error
}

func (db *TodoRepository) Create(params model.Todo) (data model.Todo, err error) {

	data = model.Todo{
		ActivityGroupId: params.ActivityGroupId,
		Title:           params.Title,
		Priority:        params.Priority,
		IsActive:        params.IsActive,
	}
	err = db.db.Debug().Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (db *TodoRepository) Update(id int, params model.Todo) (rowAffected int64, err error) {

	data := map[string]interface{}{
		"title":     params.Title,
		"priority":  params.Priority,
		"is_active": params.IsActive,
	}

	query := db.db.Debug().Model(&params)

	query = query.Where("todo_id = ?", id)

	query = query.Updates(data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, err
}
