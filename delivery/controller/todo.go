package controller

import (
	"devcode/constant"
	"devcode/delivery/common"
	"devcode/model"
	"devcode/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Todo repository.TodoRepository
}

func NewTodoController(todo *repository.TodoRepository) *TodoController {
	return &TodoController{
		Todo: *todo,
	}
}

func (rp TodoController) GetAll(ctx echo.Context) error {
	response := common.ResponseBody{}

	activityGroupId := ctx.QueryParam("activity_group_id")

	var id interface{}
	id, err := strconv.Atoi(activityGroupId)
	if err != nil {
		id = nil
	}

	data, err := rp.Todo.GetAll(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	dataMapping := []common.TodoDataResponse{}
	for _, v := range data {
		vData := common.TodoDataResponse{
			Id:              v.TodoId,
			ActivityGroupId: v.ActivityGroupId,
			Title:           v.Title,
			IsActive:        v.IsActive,
			Priority:        v.Priority,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		}

		dataMapping = append(dataMapping, vData)
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (rp TodoController) GetOne(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	data, err := rp.Todo.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	if data.TodoId == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Todo", id))
	}

	v := data
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (rp TodoController) Create(ctx echo.Context) error {

	isActive := true
	priority := constant.VeryHigh

	request := common.TodoCreate{
		IsActive: &isActive,
		Priority: &priority,
	}

	response := common.ResponseBody{}

	ctx.Bind(&request)
	// err := ctx.Bind(&request)
	// if err != nil {
	// 	data := reflect.ValueOf(request)
	// 	fieldNum := data.NumField()
	// 	reflectType := data.Type()

	// 	for i := 0; i < fieldNum; i++ {
	// 		if strings.Contains(err.Error(), strcase.SnakeCase(reflectType.Field(i).Name)) {
	// 			return ctx.JSON(http.StatusBadRequest, response.BadRequest(reflectType.Field(i).Name, reflectType.Field(i).Type.Name()))
	// 		}
	// 	}
	// }

	if err := ctx.Validate(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
			return ctx.JSON(http.StatusBadRequest, response.BadRequest(err.Field(), err.Tag()))
		}
	}

	model := model.Todo{
		ActivityGroupId: request.ActivityGroupId,
		Title:           request.Title,
		Priority:        request.Priority,
		IsActive:        request.IsActive,
	}

	// if model.IsActive == "" {
	// 	model.IsActive = constant.VeryHigh
	// }

	// if model.Priority == "" {
	// 	model.Priority = constant.VeryHigh
	// }

	data, err := rp.Todo.Create(model)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	return ctx.JSON(http.StatusCreated, response.Success(dataMapping))
}

func (rp TodoController) Delete(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	rowAffected, err := rp.Todo.Delete(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Todo", id))
	}

	dataMapping := common.DataDeleteResponse{}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (rp TodoController) Update(ctx echo.Context) error {
	request := common.TodoUpdate{}
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	data, err := rp.Todo.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}
	if data.TodoId == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Activity", id))
	}

	title := data.Title
	isActive := data.IsActive
	veryHigh := constant.VeryHigh

	request = common.TodoUpdate{
		Title:    title,
		IsActive: isActive,
		Priority: &veryHigh,
	}

	ctx.Bind(&request)
	// err := ctx.Bind(&request)
	// if err != nil {
	// 	data := reflect.ValueOf(request)
	// 	fieldNum := data.NumField()
	// 	reflectType := data.Type()

	// 	for i := 0; i < fieldNum; i++ {
	// 		if strings.Contains(err.Error(), strcase.SnakeCase(reflectType.Field(i).Name)) {
	// 			return ctx.JSON(http.StatusBadRequest, response.BadRequest(reflectType.Field(i).Name, reflectType.Field(i).Type.Name()))
	// 		}
	// 	}
	// }

	// if err := ctx.Validate(request); err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fmt.Println(err.Field(), err.Tag())
	// 		return ctx.JSON(http.StatusBadRequest, response.BadRequest(err.Field(), err.Tag()))
	// 	}
	// }

	model := model.Todo{
		Title:    request.Title,
		Priority: request.Priority,
		IsActive: request.IsActive,
	}

	_, err = rp.Todo.Update(intId, model)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	data, err = rp.Todo.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}
