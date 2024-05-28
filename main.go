package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veysoff/golang-crud-postgres-server/controllers"
	"github.com/veysoff/golang-crud-postgres-server/initializers"
	"github.com/veysoff/golang-crud-postgres-server/routes"
)

var (
	server *gin.Engine

	TaskController      controllers.TaskController
	TaskRouteController routes.TaskRouteController

	AuthentificationController      controllers.AuthentificationController
	AuthentificationRouteController routes.AuthentificationRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	TaskController = controllers.NewTaskController(initializers.DB)
	TaskRouteController = routes.NewRouteTaskController(TaskController)

	AuthentificationController = controllers.NewAuthController(initializers.DB)
	AuthentificationRouteController = routes.NewRouteAuthentificationController(AuthentificationController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Server is runing. Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	TaskRouteController.TaskRoute(router)
	AuthentificationRouteController.AuthRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
