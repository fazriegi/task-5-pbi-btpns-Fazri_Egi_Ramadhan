package router

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()

	userAuthController := controllers.UserAuthController{}
	r.POST("/users/register", userAuthController.Register)

	return r
}
