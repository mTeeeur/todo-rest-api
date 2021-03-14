package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"github.com/mTeeeur/todo-rest-api/pkg/repository"
	"time"
)

const (
	salt     = "arozaypalanalapyazora"
	tokenTTL = time.Hour * 8
	loginKey = "ylukomoryadubzelenii"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo_rest_api.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generateHashPassword(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(loginKey))
}
func (s *AuthService) ParseToken(token string) (int, error) {
	tkn, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid login method")
		}

		return []byte(loginKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	return claims.UserId, nil
}
