package controllers

import (
	"errors"
	"fmt"
	"log"
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

func GenerateToken(email string, first_name string, last_name string, user_id string) (string,string, error) {

	claims := &jwtCustomClaims{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		UserID:    user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(5 * time.Minute).Unix(),
		},
	}

	refreshClaims := &jwtCustomClaims{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		UserID:    user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "","", err
	}

	refreshTokenString, err := refreshToken.SignedString(sampleSecretKey)
	if err != nil {
		return "","" ,err
	}

	fmt.Print(tokenString)
	fmt.Print(refreshTokenString)

	return tokenString,refreshTokenString, nil

}

func VerifyToken(tokenString string) (claims *jwtCustomClaims, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return sampleSecretKey, nil
	})
	if err != nil {
		log.Println("Authorization Failed")
		return claims,errors.New("Authorization Failed....")
	}

	claims, ok := token.Claims.(*jwtCustomClaims)

	if !ok {
		log.Println("Claims Extraction Error")
		return claims,errors.New("Claims Extraction Error...")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token Expired....")
		return claims,errors.New("Token is expired")
	}

	return claims,nil
}

