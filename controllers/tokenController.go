package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type jwtCustomClaims struct {
	Email     string
	FirstName string
	LastName  string
	UserID    string
	jwt.StandardClaims
}

func GenerateToken(email string, first_name string, last_name string, user_id string) (string, error) {

	claims := &jwtCustomClaims{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		UserID:    user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(10 * time.Minute).Unix(),
		},
	}

	refreshClaims:=&jwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	refreshToken:=jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),refreshClaims)
	
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	refreshTokenString, err :=  refreshToken.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	
	fmt.Print(tokenString)
	fmt.Print(refreshTokenString)

	return tokenString, nil

}

func VerifyToken(token string) error {
	return errors.New("Test")
}
