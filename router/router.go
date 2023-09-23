package router

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()

	userAuthController := controllers.UserAuthController{}
	r.POST("/users/register", userAuthController.Register)
	r.POST("/users/login", userAuthController.Login)

	userController := controllers.UserController{}
	r.PUT("/users/:userId", userController.Update)
	r.DELETE("/users/:userId", userController.Delete)

	photoController := controllers.PhotoController{}
	r.POST("/photos", photoController.Add)

	return r
}

