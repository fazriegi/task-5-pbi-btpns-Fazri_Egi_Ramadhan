package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/helpers"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/middlewares"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"github.com/gin-gonic/gin"
)

type PhotoController struct{}

func (p *PhotoController) Add(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]
	jwtToken := strings.Split(authHeader, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	
	var photo models.Photo
	if err := c.Bind(&photo); err != nil {
		log.Println("failed to binding data: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	file, _ := c.FormFile("file")
	filePath := "user-img/" + file.Filename
	c.SaveUploadedFile(file, filePath)

	wd,err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	fmt.Println("pa",parent)

	photo.UserID = uint(userId)
	photo.PhotoURL = filePath
	photoQuery := queries.PhotoQuery{}

	if err := photoQuery.Save(&photo); err != nil {
		log.Println("failed save photo to database: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}
}