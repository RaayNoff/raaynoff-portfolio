package authServices

import (
	authDtos "core-service/internal/modules/auth/domain/dtos"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	secret []byte
}

// TODO CFG INIT
func NewAuthService() *AuthService {
	return &AuthService{
		//secret:
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

func (s *AuthService) Login(dto *authDtos.AuthLoginDto) {
	fmt.Println(dto)
}
