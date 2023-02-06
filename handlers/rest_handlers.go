package handlers

import (
	"errors"
	"jwt_athentication/controllers"
	"jwt_athentication/models"
	"jwt_athentication/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignInEndPoint(c *gin.Context) {

	user_data := models.User{}

	err := c.Bind(&user_data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = controllers.ValidateUser(&user_data)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user_data.Created_at = time.Now()
	user_data.Updated_at = time.Now()
	user_data.ID = primitive.NewObjectID()

	err = repositories.UserCreate(&user_data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Println(user_data)
	c.JSON(http.StatusAccepted, "User Signed Up Successfully")

}

func LoginEndPoint(c *gin.Context) {

	loginData := models.LoginRequest{}

	err := c.Bind(&loginData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := controllers.AuthenticateUser(&loginData)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	token_value, refresh_token_value, err1 := controllers.GenerateToken(user.Email, user.FirstName, user.LastName, user.UserId)
	log.Println(token_value)
	log.Println(refresh_token_value)
	if err1 != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = controllers.UpdateToken(token_value, refresh_token_value, user.Email)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := models.LoginResponse{
		Token: token_value,
	}
	c.SetCookie("refresh_token", refresh_token_value, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, response)

}

func RefreshToken(c *gin.Context) {

	cookie, err := c.Cookie("refresh_token")
	log.Println(cookie)

	if err != nil {
		log.Println("No refresh token found")
		c.JSON(http.StatusInternalServerError, err)
		return

	}

	claims, err := controllers.VerifyToken(cookie)

	if err != nil {
		log.Println("Refesh Token not verified")
		c.JSON(http.StatusInternalServerError, err)
		return 
	}
	token_value, refresh_token_value, err1 := controllers.GenerateToken(claims.Email, claims.FirstName, claims.LastName, claims.Id)

	if err1 != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = controllers.UpdateToken(token_value, refresh_token_value, claims.Email)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := models.LoginResponse{
		Token: token_value,
	}

	c.SetCookie("refresh_token", refresh_token_value, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, response)

}

func GetProfiles(c *gin.Context) {
	result, err := repositories.GetAll()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, result)

}

func PasswordChange(c *gin.Context)(){
	requestModel:=models.PasswordChange{}
	err:=c.BindJSON(requestModel)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}


	err=repositories.UserUpdatePassword(requestModel.Email,requestModel.NewPassword,requestModel.OldPassword)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError,errors.New("Password Change failed.."))
		return
	}
	c.JSON(http.StatusAccepted,"Password Changed Successfully")

}