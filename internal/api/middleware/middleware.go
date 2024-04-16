package middleware

import (
	"CleanOffice/config"
	"CleanOffice/internal/models"
	"CleanOffice/internal/repository"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "вы не авторизованны",
		})
		return
	}
	if err != nil {
		c.Status(http.StatusUnauthorized)
	}
	env, _ := config.LoadConfig()
	secret := env.Secret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Status(http.StatusUnauthorized)
		}
		var user models.User
		repository.DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.Status(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.Status(http.StatusUnauthorized)
	}
}
func RequireRefresh(c *gin.Context) {
	refreshTokenString, err := c.Cookie("RefreshToken")
	if refreshTokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "вы не авторизованны",
		})
		return
	}
	if err != nil {
		c.Status(http.StatusUnauthorized)
	}
	decodedRefreshToken, err := base64.StdEncoding.DecodeString(refreshTokenString)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid refresh token",
		})
	}
	rtString := string(decodedRefreshToken)
	env, _ := config.LoadConfig()
	Secret := env.Secret
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(rtString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})
	Id := claims["sub"]
	var user models.User
	repository.DB.First(&user, Id)
	accessTUSER, _ := c.MustGet("user").(models.User)
	userId := accessTUSER.Id
	if user.Id == userId {
		c.Set("user", user)
	}
}
func Cors() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}
}
