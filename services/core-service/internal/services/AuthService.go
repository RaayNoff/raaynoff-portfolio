package services

import (
	dtos "core-service/internal/dtos/auth"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	secret []byte
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{
		secret: []byte(secret),
	}
}

func (s *AuthService) GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"userId": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}

func (s *AuthService) Login(dto *dtos.AuthLoginDto) {

}
