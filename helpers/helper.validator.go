package helpers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/middlewares"

	"github.com/asaskevich/govalidator"
)

func ValidateUserInputForAuthentication(userValidation interface{}) (bool, error) {
	result, err := govalidator.ValidateStruct(userValidation)
	if err != nil {
		return result, err
	}

	return result, nil
}

func ValidateUserToken(authHeader []string) (uint, uint, error) {
	if authHeader ==  nil{
		log.Println("authorization header is not specified")
		return 0, http.StatusBadRequest, errors.New("authorization header is not specified")
	}

	authorization := authHeader[0]
	
	if authorization == "" {
		log.Println("authorization token is not specified")
		return 0, http.StatusBadRequest, errors.New("authorization token is not specified")
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := uint(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("failed to extract jwt token: ", err)
		return 0, http.StatusBadRequest, err
	}

	isUserExist, err := CheckUserId(userId)
	if !isUserExist && err != nil {
		log.Println("signature is invalid: ", err)
		return 0, http.StatusNotFound, errors.New("signature is invalid")
	}

	return userId, http.StatusOK, nil
}