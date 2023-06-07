package service

import (
	. "github.com/ksi2021/go-kata/module4/webserver/http/homework/repository/user"
)

type TodoService interface {
	ListUsers() ([]User, error)
	CreateUser(username, password, email string) error
	GetUser(id int) (User, error)
}

type todoService struct {
	repository UserRepository
}

func MakeService(path string) *todoService {
	return &todoService{repository: &FileUserRepository{FilePath: path}}
}

func (s *todoService) ListUsers() ([]User, error) {
	users, err := s.repository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *todoService) CreateUser(username, password, email string) bool {
	user := User{UserName: username, Password: password, Email: email}
	if _, err := s.repository.CreateUser(user); err != nil {
		return false
	}

	return true
}

func (s *todoService) GetUser(id int) (User, error) {
	user, err := s.repository.GetUser(id)

	if err != nil {
		return User{}, err
	}
	return user, nil
}
