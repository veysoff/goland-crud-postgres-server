package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/veysoff/golang-crud-postgres-server/controllers"
)

type AuthentificationRouteController struct {
	authController controllers.AuthentificationController
}

func NewRouteAuthentificationController(authController controllers.AuthentificationController) AuthentificationRouteController {
	return AuthentificationRouteController{authController}
}

func (rc *AuthentificationRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
}
