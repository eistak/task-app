package routes

import (
	"go-ws/task-app/pkg"

	"github.com/labstack/echo/v4"
)

func Route(group *echo.Group) {
	group.GET("/tasks", pkg.GetTasks())
	group.GET("/tasks/:id", pkg.GetTask())
	group.POST("/tasks", pkg.AddTask())
	group.PUT("/tasks/:id", pkg.UpdateTask())
	group.DELETE("/tasks/:id", pkg.RemoveTask())
}
