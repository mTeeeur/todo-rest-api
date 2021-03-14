package service

import (
	"crypto/sha1"
	"fmt"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"github.com/mTeeeur/todo-rest-api/pkg/repository"
)

const salt = "arozaypalanalapyazora"

type AuthService struct {
	repo repository.Authorization
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
