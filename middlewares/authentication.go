package middlewares

import (
	"jwt_athentication/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context)  {

	log.Println("Authetication middleware....")

	client_token:=c.Request.Header.Get("token")

	if client_token==" "{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"No Bearer Token Provided",
		})
		c.Abort()
	}

	claims,err:=controllers.VerifyToken(client_token)
	
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Authorization Failed for the user",
		})
		c.Abort()
	}
	c.Set("email",claims.Email)
	c.Set("first_name",claims.FirstName)
	c.Set("last_name",claims.LastName)
	log.Println("Authorization Successful.....")
	c.Next()

	
}