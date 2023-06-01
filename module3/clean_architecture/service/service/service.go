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

func (s *todoService) ListTodos() ([]Task, error) {
	tasks, err := s.repository.GetTasks()

	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (s *todoService) CreateTodo(title string) error {
	todo := Task{Title: title}
	if _, err := s.repository.CreateTask(todo); err != nil {
		return err
	}

	return nil
}

func (s *todoService) CompleteTodo(todo Task) error {
	todo.Status = true
	if _, err := s.repository.UpdateTask(todo); err != nil {
		return err
	}

	return nil
}

func (s *todoService) RemoveTodo(todo Task) error {
	if err := s.repository.DeleteTask(todo.ID); err != nil {
		return err
	}
	return nil
}
