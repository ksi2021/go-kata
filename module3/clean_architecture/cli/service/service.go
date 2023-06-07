package service

import . "github.com/ksi2021/go-kata/module3/clean_architecture/service/repo"

type TodoService interface {
	ListTodos() ([]Task, error)
	CreateTodo(title string) error
	CompleteTodo(todo Task) error
	RemoveTodo(todo Task) error
}

type todoService struct {
	repository TaskRepository
}

func MakeService(path string) *todoService {
	return &todoService{repository: &FileTaskRepository{FilePath: path}}
}

func (s *todoService) ListTodos() ([]Task, bool) {
	tasks, err := s.repository.GetTasks()

	if err != nil {
		return nil, false
	}

	return tasks, true
}

func (s *todoService) CreateTodo(title string) bool {
	todo := Task{Title: title}
	if _, err := s.repository.CreateTask(todo); err != nil {
		return false
	}

	return true
}

func (s *todoService) CompleteTodo(todo Task) bool {
	todo.Status = true
	if _, err := s.repository.UpdateTask(todo); err != nil {
		return false
	}

	return true
}

func (s *todoService) RemoveTodo(todo Task) bool {
	if err := s.repository.DeleteTask(todo.ID); err != nil {
		return false
	}
	return true
}
