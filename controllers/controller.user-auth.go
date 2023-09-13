package controllers

import (
	"log"
	"net/http"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
)

type UserAuthController struct{}

func (ua *UserAuthController) Register(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Println("failed to binding data: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	emailRegistered, err := helpers.IsRegistered(user.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if emailRegistered {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email already registered",
		})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Println("failed hashing password: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed register account",
		})
		return
	}

	user.Password = hashedPassword

	if err = queries.Save(&user); err != nil {
		log.Println("failed save user to database: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed register account",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success register account",
	})
}
