package controller

import (
	"devcode/constant"
	"devcode/delivery/common"
	"devcode/model"
	"devcode/repository"
	"net/http"
	"strconv"

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

func (pc TodoController) GetAll(ctx echo.Context) error {
	response := common.ResponseBody{}

	activityGroupId := ctx.QueryParam("activity_group_id")

	var id interface{}
	id, err := strconv.Atoi(activityGroupId)
	if err != nil {
		id = nil
	}

	data, err := pc.Todo.GetAll(id)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	dataMapping := []common.TodoDataResponse{}
	for _, v := range data {
		vData := common.TodoDataResponse{
			Id:              v.TodoId,
			ActivityGroupId: v.ActivityGroupId,
			Title:           v.Title,
			IsActive:        v.IsActive,
			Priority:        v.Priority,
			CreatedAt:       *v.CreatedAt,
			UpdatedAt:       *v.UpdatedAt,
		}

		dataMapping = append(dataMapping, vData)
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (pc TodoController) GetOne(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	data, err := pc.Todo.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	if data.TodoId == 0 {
		return ctx.JSON(http.StatusOK, response.NotFound("Todo", id))
	}

	v := data
	// dataMapping := common.TodoDataResponse{}
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       *v.CreatedAt,
		UpdatedAt:       *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (pc TodoController) Create(ctx echo.Context) error {
	response := common.ResponseBody{}

	request := common.TodoCreate{}

	ctx.Bind(&request)

	model := model.Todo{
		ActivityGroupId: request.ActivityGroupId,
		Title:           request.Title,
		Priority:        constant.VeryHigh,
		IsActive:        request.IsActive,
	}

	data, err := pc.Todo.Create(model)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       *v.CreatedAt,
		UpdatedAt:       *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (pc TodoController) Delete(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	rowAffected, err := pc.Todo.Delete(intId)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusOK, response.NotFound("Todo", id))
	}

	// if data.TodoId == 0 {
	// 	return ctx.JSON(http.StatusOK, response.NotFound("Todo", id))
	// }

	return ctx.JSON(http.StatusOK, response.Success(nil))
}

func (pc TodoController) Update(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	request := common.TodoUpdate{}
	ctx.Bind(&request)

	model := model.Todo{
		Title:    request.Title,
		Priority: request.Priority,
		IsActive: request.IsActive,
	}

	rowAffected, err := pc.Todo.Update(intId, model)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusOK, response.NotFound("Todo", id))
	}

	data, err := pc.Todo.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusOK, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.TodoDataResponse{
		Id:              v.TodoId,
		ActivityGroupId: v.ActivityGroupId,
		Title:           v.Title,
		IsActive:        v.IsActive,
		Priority:        v.Priority,
		CreatedAt:       *v.CreatedAt,
		UpdatedAt:       *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}
