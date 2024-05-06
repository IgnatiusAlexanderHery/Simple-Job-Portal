package auth

import (
	"Simple-Job-Portal/model"
	"Simple-Job-Portal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func GetJWTKey() []byte {
	return jwtKey
}

func AuthMiddlewareTalent() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			utils.UnauthorizedError(c, "Unauthorized token")
			return
		}

		tokenString := cookie
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid || claims.Role != "talent" {
			utils.UnauthorizedError(c, "Unauthorized role")
			return
		}

		csrfToken, csrfErr := c.Cookie("X-CSRF-Token")

		if csrfErr != nil || csrfToken == "" {
			utils.UnauthorizedError(c, "Unauthorized csrf")
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func AuthMiddlewareEmployer() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			utils.UnauthorizedError(c, "Unauthorized")
			return
		}

		tokenString := cookie
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid || claims.Role != "employer" {
			utils.UnauthorizedError(c, "Unauthorized")
			return
		}

		csrfToken, csrfErr := c.Cookie("X-CSRF-Token")
		if csrfErr != nil || csrfToken == "" {
			utils.UnauthorizedError(c, "Unauthorized")
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
