package controllers

import (
	"log"
	"net/http"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/middlewares"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
)

type UserAuthController struct{}

func (ua *UserAuthController) Register(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	emailRegistered, err := helpers.IsRegistered(user.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if emailRegistered {
		helpers.SendResponse(c, http.StatusBadRequest, "email already registered", nil)

		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Println("failed hashing password: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "failed register account", nil)

		return
	}

	user.Password = hashedPassword

	if err = queries.Save(&user); err != nil {
		log.Println("failed save user to database: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "failed register account", nil)

		return
	}

	helpers.SendResponse(c, http.StatusOK, "success register account", nil)
}

func (ua *UserAuthController) Login(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	emailRegistered, err := helpers.IsRegistered(user.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if !emailRegistered {
		helpers.SendResponse(c, http.StatusUnauthorized, "email or password wrong!", nil)

		return
	}

	userData, err := queries.GetUser(user.Email)

	if err != nil {
		log.Println("failed get user hashed password: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := helpers.ComparePassword(user.Password, userData.Password); err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, "email or password wrong!", nil)

		return
	}

	jwtToken, err := middlewares.CreateJWTToken(userData.ID)

	data := map[string]string{
		"token": jwtToken,
	}

	helpers.SendResponse(c, http.StatusOK, "login success", data)
}
