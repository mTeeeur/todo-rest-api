package service

import (
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"github.com/mTeeeur/todo-rest-api/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo_rest_api.User) (int, error) {
	return s.repo.CreateUser(user)
}
