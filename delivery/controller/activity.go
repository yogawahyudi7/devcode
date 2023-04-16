package controller

import (
	"devcode/delivery/common"
	"devcode/model"
	"devcode/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	Activity repository.ActivityRepository
}

func NewActivityController(activity *repository.ActivityRepository) *ActivityController {
	return &ActivityController{
		Activity: *activity,
	}
}

func (rp ActivityController) GetAll(ctx echo.Context) error {
	response := common.ResponseBody{}

	data, err := rp.Activity.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	dataMapping := []common.ActivityDataResponse{}
	for _, v := range data {
		vData := common.ActivityDataResponse{
			Id:        v.ActivityId,
			Title:     v.Title,
			Email:     v.Email,
			CreatedAt: *v.CreatedAt,
			UpdatedAt: *v.UpdatedAt,
		}

		dataMapping = append(dataMapping, vData)
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (rp ActivityController) GetOne(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	data, err := rp.Activity.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	if data.ActivityId == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Activity", id))
	}

	v := data
	dataMapping := common.ActivityDataResponse{
		Id:        v.ActivityId,
		Title:     v.Title,
		Email:     v.Email,
		CreatedAt: *v.CreatedAt,
		UpdatedAt: *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}

func (rp ActivityController) Create(ctx echo.Context) error {
	request := common.ActivityCreate{}
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

	model := model.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	data, err := rp.Activity.Create(model)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.ActivityDataResponse{
		Id:        v.ActivityId,
		Title:     v.Title,
		Email:     v.Email,
		CreatedAt: *v.CreatedAt,
		UpdatedAt: *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusCreated, response.Success(dataMapping))
}

func (rp ActivityController) Delete(ctx echo.Context) error {
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	rowAffected, err := rp.Activity.Delete(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Activity", id))
	}

	return ctx.JSON(http.StatusOK, response.Success(nil))
}

func (rp ActivityController) Update(ctx echo.Context) error {
	request := common.ActivityUpdate{}
	response := common.ResponseBody{}

	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

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

	model := model.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	rowAffected, err := rp.Activity.Update(intId, model)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound("Activity", id))
	}

	data, err := rp.Activity.GetOne(intId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err))
	}

	v := data
	dataMapping := common.ActivityDataResponse{
		Id:        v.ActivityId,
		Title:     v.Title,
		Email:     v.Email,
		CreatedAt: *v.CreatedAt,
		UpdatedAt: *v.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response.Success(dataMapping))
}
