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

func (pc *TaskRouteController) TaskRoute(rg *gin.RouterGroup) {

	router := rg.Group("tasks")
	// TODO add such middleware
	// router.Use(middleware.DeserializeUser())
	router.POST("/", pc.taskController.CreateTask)
	router.GET("/", pc.taskController.FindTasks)
	router.PUT("/", pc.taskController.UpdateTask)
	router.GET("/:taskId", pc.taskController.FindTaskById)
	router.DELETE("/", pc.taskController.DeleteTask)
}
