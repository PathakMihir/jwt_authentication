package handlers

import (
	"jwt_athentication/controllers"
	"jwt_athentication/models"
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

	err = controllers.InsertUser(&user_data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return 
	}
	log.Println(user_data)
	c.JSON(http.StatusAccepted, "User Signed Up Successfully")

}

func LoginEndPoint(c *gin.Context) {

	loginData:=models.LoginRequest{}

	err:=c.Bind(&loginData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user,err:=controllers.AuthenticateUser(&loginData)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	token_value,err1:=controllers.GenerateToken(user.Email,user.FirstName,user.LastName,user.UserId)

	if err1 != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err=controllers.UpdateToken(token_value ,&user)
	
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response:=models.LoginResponse{
		Token: token_value,
	}
	

	c.JSON(http.StatusOK,response)


}

func GetProfiles(c *gin.Context) {

}
