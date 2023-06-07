package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Repository layer

// User represents a to-do task
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRepository is a repository interface for tasks
type UserRepository interface {
	GetUsers() ([]User, error)
	GetUser(id int) (User, error)
	CreateUser(user User) (User, error)
}

// FileUserRepository is a file-based implementation of TaskRepository
type FileUserRepository struct {
	FilePath string
}

// GetUsers returns all tasks from the repository
func (repo *FileUserRepository) GetUsers() ([]User, error) {
	var users []User

	file, err := os.Open(repo.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(content, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *FileUserRepository) SaveUsers(users []User) error {
	file, err := os.Create(repo.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(repo.FilePath, data, 0666)
	if err != nil {
		panic(err)
	}

	return nil
}

// GetUser returns a single task by its ID
func (repo *FileUserRepository) GetUser(id int) (User, error) {
	var user User

	tasks, err := repo.GetUsers()
	if err != nil {
		return user, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return user, fmt.Errorf("user not found")
}

// CreateUser adds a new task to the repository
func (repo *FileUserRepository) CreateUser(user User) (User, error) {
	users, err := repo.GetUsers()
	if err != nil {
		return user, err
	}

	user.ID = len(users) + 1
	users = append(users, user)

	if err := repo.SaveUsers(users); err != nil {
		return user, err
	}

	return user, nil
}
