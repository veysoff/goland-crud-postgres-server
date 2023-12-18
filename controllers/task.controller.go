package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/veysoff/golang-crud-postgres-server/models"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

func NewTaskController(DB *gorm.DB) TaskController {
	return TaskController{DB}
}

// [...] Create Handler
func (tc *TaskController) CreateTask(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newTask := models.Task{
		Title:       payload.Title,
		Description: payload.Description,
		//User:        currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := tc.DB.Create(&newTask)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Task with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTask})
}

// [...] Update Handler
func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	//currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateTask
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedTask models.Task
	result := tc.DB.First(&updatedTask, "id = ?", taskId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	taskToUpdate := models.Task{
		Title:       payload.Title,
		Description: payload.Description,
		IsDone:      payload.IsDone,
		//User:        currentUser.ID,
		CreatedAt: updatedTask.CreatedAt,
		UpdatedAt: now,
	}

	tc.DB.Model(&updatedTask).Updates(taskToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTask})
}

func (pc *TaskController) FindTaskById(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var task models.Task
	result := pc.DB.First(&task, "id = ?", taskId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": task})
}

func (tc *TaskController) FindTasks(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tasks []models.Task
	results := tc.DB.Limit(intLimit).Offset(offset).Find(&tasks)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tasks), "data": tasks})
}

func (pc *TaskController) DeleteTask(ctx *gin.Context) {
	postId := ctx.Param("taskId")

	result := pc.DB.Delete(&models.Task{}, "id = ?", postId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No task with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
