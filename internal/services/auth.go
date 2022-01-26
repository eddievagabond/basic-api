package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/eddievagabond/internal/models"
	"github.com/eddievagabond/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	Authenticate(c context.Context, email, password string) (*models.User, error)
	Register(c context.Context, u *models.CreateUserParams) (*models.User, error)
	GenerateAccessToken(userId string) (string, error)
	GenerateRefreshToken(userId string) (string, error)
	GenerateCustomKey(userID string, tokenHash string) string
	ValidateAccessToken(tokenString string) (string, error)
	ValidateRefreshToken(tokenString string) (string, error)
}

type AccessTokenClaims struct {
	UserId  string `json:"user_id"`
	KeyType string `json:"key_type"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserId  string `json:"user_id"`
	KeyType string `json:"key_type"`
	jwt.StandardClaims
}

type authService struct {
	config   *util.Configuration
	authRepo models.AuthRepository
}

func NewAuthService(config *util.Configuration, authRepo models.AuthRepository) *authService {
	return &authService{
		config:   config,
		authRepo: authRepo,
	}
}

// Authenticate a user by email and password
func (a *authService) Authenticate(c context.Context, email, password string) (*models.User, error) {
	user, err := a.authRepo.GetByEmail(c, email)
	if err != nil {
		return nil, err
	}

	err = a.comparePasswordHash(password, user.HashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register a new user
func (a *authService) Register(c context.Context, u *models.CreateUserParams) (*models.User, error) {
	hashedPassword, err := a.hashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	user, err := a.authRepo.Create(c, u, hashedPassword)
	if err != nil {
		return nil, err
	}

	return user.Sanitize(), nil
}

// GenerateAccessToken generates an access token for a user
func (a *authService) GenerateAccessToken(userId string) (string, error) {
	tokenType := "access"

	claims := AccessTokenClaims{
		userId,
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(a.config.JwtExpiration)).Unix(),
			Issuer:    "basic-api.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(a.config.AccessTokenPrivateKeyPath)

	if err != nil {
		return "", fmt.Errorf("error reading private key: %s", err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateRefreshToken generates a refresh token for a user
func (a *authService) GenerateRefreshToken(userId string) (string, error) {
	tokenType := "refresh"

	claims := RefreshTokenClaims{
		userId,
		tokenType,
		jwt.StandardClaims{
			Issuer: "basic-api.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(a.config.RefreshTokenPrivateKeyPath)
	if err != nil {
		return "", fmt.Errorf("error reading private key: %s", err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateCustomKey generates a custom key for a user
func (a *authService) GenerateCustomKey(userID string, tokenHash string) string {
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// ValidateAccessToken validates an access token
func (a *authService) ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("can't handle %v", token.Header["alg"])
		}
		verifyBytes, err := ioutil.ReadFile(a.config.AccessTokenPublicKeyPath)
		if err != nil {
			return nil, fmt.Errorf("error reading public key: %s", err)
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing public key: %s", err)
		}

		return verifyKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %s", err)
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid || claims.UserId == "" || claims.KeyType != "access" {
		return "", fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

// ValidateRefreshToken validates a refresh token
func (a *authService) ValidateRefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("can't handle %v", token.Header["alg"])
		}
		verifyBytes, err := ioutil.ReadFile(a.config.RefreshTokenPublicKeyPath)
		if err != nil {
			return nil, fmt.Errorf("error reading public key: %s", err)
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing public key: %s", err)
		}

		return verifyKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %s", err)
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid || claims.UserId == "" || claims.KeyType != "refresh" {
		return "", fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

func (a *authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (a *authService) comparePasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
