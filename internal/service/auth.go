package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/eddievagabond/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type AccessTokenClaims struct {
	UserID  string `json:"user_id"`
	KeyType string `json:"key_type"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserID    string `json:"user_id"`
	CustomKey string `json:"custom_key"`
	KeyType   string `json:"key_type"`
	jwt.StandardClaims
}

type AuthService struct {
	config *util.Configuration
}

func NewAuthService(config *util.Configuration) *AuthService {
	return &AuthService{
		config: config,
	}
}

func (auth *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (auth *AuthService) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
