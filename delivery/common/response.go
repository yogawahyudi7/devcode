package common

import (
	"fmt"
	"time"

	"github.com/stoewer/go-strcase"
)

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ActivityDataResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoDataResponse struct {
	Id              uint      `json:"id"`
	ActivityGroupId int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (r ResponseBody) Success(data interface{}) ResponseBody {

	return ResponseBody{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
}

func (r ResponseBody) BadRequest(object string, tag string) ResponseBody {

	object = strcase.SnakeCase(object)
	message := ""

	switch tag {
	case "email":
		message = fmt.Sprintf("%v is invalid. ex : xample@mail.com", object)
	case "required":
		message = fmt.Sprintf("%v cannot be null", object)
	default:
		message = fmt.Sprintf("%v value type must be %v", object, tag)
	}
	return ResponseBody{
		Status:  "Bad Request",
		Message: message,
	}
}

func (r ResponseBody) NotFound(object string, id string) ResponseBody {

	message := fmt.Sprintf("%v with ID %v Not Found", object, id)
	responsebody := ResponseBody{
		Status:  "Not Found",
		Message: message,
	}

	return responsebody
}

func (r ResponseBody) InternalServerError(err error) ResponseBody {

	message := err.Error()
	responsebody := ResponseBody{
		Status:  "Internal Server Error",
		Message: message,
	}

	return responsebody
}
