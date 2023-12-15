package main

import (
	"fmt"
	"log"

	"github.com/veysoff/golang-crud-postgres-server/initializers"
	"github.com/veysoff/golang-crud-postgres-server/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Task{})
	fmt.Println("👍 Migration complete")
}
