package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/veysoff/golang-crud-postgres-server/controllers"
)

type TaskRouteController struct {
	taskController controllers.TaskController
}

func NewRouteTaskController(taskController controllers.TaskController) TaskRouteController {
	return TaskRouteController{taskController}
}

func (tc *TaskRouteController) TaskRoute(rg *gin.RouterGroup) {

	router := rg.Group("tasks")
	// TODO add such middleware
	// router.Use(middleware.DeserializeUser())
	router.POST("/", tc.taskController.CreateTask)
	router.GET("/", tc.taskController.FindTasks)
	router.PUT("/", tc.taskController.UpdateTask)
	router.GET("/:taskId", tc.taskController.FindTaskById)
	router.DELETE("/", tc.taskController.DeleteTask)
}
