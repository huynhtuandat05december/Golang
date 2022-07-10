package middlewares

import (
	"golang_api/helpers"
	"golang_api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtService services.JWTService = services.NewJWTService()
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader, "ACCESS_SECRET")
		if !token.Valid {
			log.Println(token, err)
			response := helpers.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[userID]: ", claims["userID"])
		c.Set("userID", claims["userID"])
		c.Next()

	}
}
