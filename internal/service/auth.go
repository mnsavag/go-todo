package service

import (
	"crypto/sha1"
	"fmt"
	"goTodo/internal/model"
	"goTodo/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "h3;21ei0dais021h"
	signingKey = "asmnvbzixurhiewqur1"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetRegisteredUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{now.Add(tokenTTL)},
			IssuedAt:  &jwt.NumericDate{now},
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
