package main

import (
	"devcode/config"
	"devcode/constant"
	"devcode/delivery/controller"
	"devcode/delivery/middleware"
	"devcode/delivery/route"
	"devcode/repository"
	"devcode/util"

	"github.com/labstack/echo/v4"
	middlewares "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	middleware.LogMiddleware(e)
	e.Pre(middlewares.RemoveTrailingSlash())

	config := config.Get()
	db := util.InitDB(config)
	util.InitialMigrate(config, db)

	todoRepository := repository.NewTodoRepository(db)
	activityRepository := repository.NewActivityRepository(db)

	todoController := controller.NewTodoController(todoRepository)
	activityController := controller.NewActivityController(activityRepository)

	route.RegisterPathTodo(e, todoController)
	route.RegisterPathActivity(e, activityController)

	e.Logger.Fatal(e.Start(":" + constant.Port))
}
