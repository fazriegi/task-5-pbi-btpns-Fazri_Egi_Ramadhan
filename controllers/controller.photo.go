package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/response"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/middlewares"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct{}
var photoQuery queries.PhotoQuery

func (p *PhotoController) Add(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	
	if authHeader ==  nil{
		log.Println("authorization header is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization header is not specified", nil)

		return
	}

	authorization := authHeader[0]
	
	if authorization == "" {
		log.Println("authorization token is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization token is not specified", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	
	var photo models.Photo
	if err := c.Bind(&photo); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	file, err := c.FormFile("file")
	
	if err != nil {
		log.Println("failed to binding file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	filePath, err := helpers.SaveFileToDir(c, file) 

	if err != nil {
		log.Println("failed to save file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	photo.UserID = uint(userId)
	photo.PhotoURL = filePath

	if err := photoQuery.Save(&photo); err != nil {
		log.Println("failed save photo to database: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	data := response.Photo{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID: photo.UserID,
	}

	helpers.SendResponse(c, http.StatusCreated, "success upload photo", data)
}

func (p *PhotoController) Get(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	
	if authHeader ==  nil{
		log.Println("authorization header is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization header is not specified", nil)

		return
	}

	authorization := authHeader[0]
	
	if authorization == "" {
		log.Println("authorization token is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization token is not specified", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	photo, err := photoQuery.Get(uint(userId))

	if err == gorm.ErrRecordNotFound {
		log.Println("user's photo not found: ")
		helpers.SendResponse(c, http.StatusNotFound, "user's photo not found", nil)
		
		return
	} else if err != nil {
		log.Println("failed to get user's photo: ", err)
		helpers.SendResponse(c, http.StatusNotFound, err.Error(), nil)
		
		return
	} 

	data := response.Photo{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID: photo.UserID,
	}

	helpers.SendResponse(c, http.StatusOK, "success get user's photo", data)
}

func (p *PhotoController) Update(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	
	if authHeader ==  nil{
		log.Println("authorization header is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization header is not specified", nil)

		return
	}

	authorization := authHeader[0]
	
	if authorization == "" {
		log.Println("authorization token is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization token is not specified", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photoFromDatabase, err := photoQuery.Get(uint(userId))

	if err != nil {
		log.Println("failed to get user's photo id: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	if userId != int(photoFromDatabase.UserID) || photoId != int(photoFromDatabase.ID) {
		log.Println("unauthorized to update photo")
		helpers.SendResponse(c, http.StatusUnauthorized, "can't update photo", nil)

		return
	}

	var photo models.Photo

	if err := c.Bind(&photo); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	file, err := c.FormFile("file")
	
	if err != nil {
		log.Println("failed to binding file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	filePath, err := helpers.SaveFileToDir(c, file) 

	if err != nil {
		log.Println("failed to save file: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := helpers.RemoveFileFromDir(photoFromDatabase.PhotoURL); err != nil {
		log.Println("failed to delete file from dir: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	photo.ID = uint(photoId)
	photo.PhotoURL = filePath

	if err := photoQuery.Update(&photo); err != nil {
		log.Println("failed to update user's photo: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	data := response.Photo{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID: uint(userId),
	}
	helpers.SendResponse(c, http.StatusOK, "success update user's photo", data)
}

func (p *PhotoController) Delete(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]
	
	if authHeader ==  nil{
		log.Println("authorization header is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization header is not specified", nil)

		return
	}

	authorization := authHeader[0]
	
	if authorization == "" {
		log.Println("authorization token is not specified")
		helpers.SendResponse(c, http.StatusBadRequest, "authorization token is not specified", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photoFromDatabase, err := photoQuery.Get(uint(userId))

	if err != nil {
		log.Println("failed to get user's photo id: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	if userId != int(photoFromDatabase.UserID) || photoId != int(photoFromDatabase.ID){
		log.Println("unauthorized to delete photo")
		helpers.SendResponse(c, http.StatusUnauthorized, "can't delete photo", nil)

		return
	}

	if err := helpers.RemoveFileFromDir(photoFromDatabase.PhotoURL); err != nil {
		log.Println("failed to delete file from dir: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := photoQuery.Delete(uint(photoId)); err != nil {
		log.Println("failed to delete photo from database")
		helpers.SendResponse(c, http.StatusUnauthorized, "failed to delete photo", nil)

		return
	}

	helpers.SendResponse(c, http.StatusOK, "success delete photo", nil)
}