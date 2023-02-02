package server

import (
	"jwt_athentication/handlers"
	"jwt_athentication/middlewares"

	"github.com/gin-gonic/gin"
)

func Init()  {
	
	router := gin.Default()
	gin.Logger()

	v1 := router.Group("/v1")
  {

    v1.POST("/login",  handlers.LoginEndPoint)
    v1.POST("/signIn", handlers.SignInEndPoint)

	authorizationGroup:=v1.Group("")
	authorizationGroup.Use(middlewares.Authenticate)
	authorizationGroup.GET("/user",handlers.GetProfiles)
	
  }
  	
	
	router.Run("localhost:8080")
	
}