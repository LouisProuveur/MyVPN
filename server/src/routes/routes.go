package routes

import (
	ctrl "MyVPN/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine) error {

	auth, err := ctrl.NewAuthenticator()

	if err != nil {
		return err
	}

	router.POST("/api/sign", auth.SignIn)

	router.Run()

	return nil
}
