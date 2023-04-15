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

func (td *TodoRepository) GetAll(id interface{}) (data []model.Todo, err error) {

	if id != nil {
		err = td.db.Debug().Where("activity_group_id = ?", id).Find(&data).Error
	} else {
		err = td.db.Debug().Find(&data).Error
	}

	if err != nil {
		return data, err
	}

	return data, err
}

func (td *TodoRepository) GetOne(id int) (data model.Todo, err error) {

	err = td.db.Debug().Where("todo_id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (td *TodoRepository) Delete(id int) (rowAffected int64, err error) {

	data := model.Todo{}

	query := td.db.Debug().Where("todo_id = ?", id).Delete(&data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, query.Error
}

func (td *TodoRepository) Create(params model.Todo) (data model.Todo, err error) {

	data = model.Todo{
		ActivityGroupId: params.ActivityGroupId,
		Title:           params.Title,
		Priority:        params.Priority,
		IsActive:        params.IsActive,
	}
	err = td.db.Debug().Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (td *TodoRepository) Update(id int, params model.Todo) (rowAffected int64, err error) {

	data := map[string]interface{}{
		"title":     params.Title,
		"priority":  params.Priority,
		"is_active": params.IsActive,
	}

	query := td.db.Debug().Model(&params).Omit("activity_group_id")

	query = query.Where("todo_id = ?", id)

	query = query.Updates(data)

	if query.Error != nil {
		return rowAffected, query.Error
	}

	rowAffected = query.RowsAffected

	return rowAffected, err
}
