package route

import (
	"devcode/delivery/controller"

	"github.com/labstack/echo/v4"
)

func RegisterPathActivity(e *echo.Echo, activity *controller.ActivityController) {
	e.GET("/activity-groups", activity.GetAll)
	e.GET("/activity-groups/:id", activity.GetOne)
	e.POST("/activity-groups", activity.Create)
	e.DELETE("/activity-groups/:id", activity.Delete)
	e.PATCH("/activity-groups/:id", activity.Update)
}
