package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateAccessToken(userID string) string
	GenerateRefreshToken(userID string) string
	ValidateToken(tokenString string, secretKey string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (j *jwtService) GenerateAccessToken(UserID string) string {
	accessTokenSecret := os.Getenv("ACCESS_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": UserID,
		"exp":    time.Now().AddDate(1, 0, 0).Unix(),
	})
	tokenResult, err := token.SignedString([]byte(accessTokenSecret))
	if err != nil {
		panic(err)
	}
	return tokenResult

}

func (j *jwtService) GenerateRefreshToken(UserID string) string {
	refreshTokenSecret := os.Getenv("REFRESH_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": UserID,
		"exp":    time.Now().AddDate(10, 0, 0).Unix(),
	})
	tokenResult, err := token.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		panic(err)
	}
	return tokenResult

}

func (j *jwtService) ValidateToken(tokenString string, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv(secretKey)
		fmt.Println("valid")
		return []byte(secret), nil
	})

}
