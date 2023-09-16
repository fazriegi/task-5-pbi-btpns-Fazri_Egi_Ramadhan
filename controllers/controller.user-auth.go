package controllers

import (
	"log"
	"net/http"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/app"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/middlewares"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
)

type UserAuthController struct{}

func (ua *UserAuthController) Register(c *gin.Context) {
	var userInput app.RegisterValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	isValidated, err := helpers.ValidateUserInputForAuthentication(userInput)

	if err != nil || !isValidated {
		log.Println("validation error: " + err.Error())
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	emailRegistered, err := helpers.IsRegistered(userInput.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if emailRegistered {
		helpers.SendResponse(c, http.StatusBadRequest, "email already registered", nil)

		return
	}

	hashedPassword, err := helpers.HashPassword(userInput.Password)

	if err != nil {
		log.Println("failed hashing password: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "failed register account", nil)

		return
	}

	userForDatabase := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	if err = queries.Save(&userForDatabase); err != nil {
		log.Println("failed save user to database: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "failed register account", nil)

		return
	}

	helpers.SendResponse(c, http.StatusOK, "success register account", nil)
}

func (ua *UserAuthController) Login(c *gin.Context) {
	var userInput app.LoginValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	isValidated, err := helpers.ValidateUserInputForAuthentication(userInput)

	if err != nil || !isValidated {
		log.Println("validation error: " + err.Error())
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	emailRegistered, err := helpers.IsRegistered(userInput.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if !emailRegistered {
		helpers.SendResponse(c, http.StatusUnauthorized, "email or password wrong!", nil)

		return
	}

	userData, err := queries.GetUser(userInput.Email)

	if err != nil {
		log.Println("failed get user hashed password: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := helpers.ComparePassword(userInput.Password, userData.Password); err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, "email or password wrong!", nil)

		return
	}

	jwtToken, err := middlewares.CreateJWTToken(userData.ID)

	data := map[string]string{
		"token": jwtToken,
	}

	helpers.SendResponse(c, http.StatusOK, "login success", data)
}
