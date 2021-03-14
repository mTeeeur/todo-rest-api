package service

import (
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"github.com/mTeeeur/todo-rest-api/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo_rest_api.User) (int, error) {
	return s.repo.CreateUser(user)
}
