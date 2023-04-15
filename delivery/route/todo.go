package route

import (
	"devcode/delivery/controller"

	"github.com/labstack/echo/v4"
)

func RegisterPathTodo(e *echo.Echo, todo *controller.TodoController) {
	e.GET("/todo-items", todo.GetAll)
	e.GET("/todo-items/:id", todo.GetOne)
	e.POST("/todo-items", todo.Create)
	e.DELETE("/todo-items/:id", todo.Delete)
	e.PATCH("/todo-items/:id", todo.Update)
}
