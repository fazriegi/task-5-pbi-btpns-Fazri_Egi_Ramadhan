package controllers

import (
	"log"
	"net/http"
	"strconv"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/app"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Update(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	userId, httpStatus, err := helpers.ValidateUserToken(authHeader)
	
	if err != nil {
		helpers.SendResponse(c, int(httpStatus), err.Error(), nil)
		return
	}

	var userInput app.UserValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	userIdFromParam, _ := strconv.Atoi(c.Param("userId"))

	if userId != uint(userIdFromParam) {
		helpers.SendResponse(c, http.StatusUnauthorized, "can't update user with given id", nil)
		return
	}

	emailRegistered, emailId, err := helpers.IsRegistered(userInput.Email)

	if err != nil {
		log.Println("failed to check email: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}
	
	if emailRegistered && emailId != uint(userId) {
		helpers.SendResponse(c, http.StatusBadRequest, "email already registered", nil)

		return
	}

	userForUpdate := models.User{
		Username: userInput.Username,
		Email: userInput.Email,
	}
	userForUpdate.ID = uint(userId)
	
	if err := user.Update(&userForUpdate); err != nil {
		log.Println("failed to update user: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "success update user", nil)
}

func (uc *UserController) Delete(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	userId, httpStatus, err := helpers.ValidateUserToken(authHeader)
	
	if err != nil {
		helpers.SendResponse(c, int(httpStatus), err.Error(), nil)
		return
	}

	userIdFromParam, _ := strconv.Atoi(c.Param("userId"))

	if userId == uint(userIdFromParam) {
		helpers.SendResponse(c, http.StatusUnauthorized, "can't delete user with given id", nil)
		return
	}

	if err := user.Delete(uint(userId)); err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	helpers.SendResponse(c, http.StatusOK, "success delete user", nil)
}