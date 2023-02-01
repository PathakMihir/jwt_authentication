package server

import (
	"jwt_athentication/handlers"

	"github.com/gin-gonic/gin"
)

func Init()  {
	
	router := gin.Default()
	gin.Logger()

	v1 := router.Group("/v1")
  {
    v1.POST("/login",  handlers.LoginEndPoint)
    v1.POST("/signUp", handlers.SignUpEndPoint)

  }
	
	router.Run("localhost:8080")
	
}