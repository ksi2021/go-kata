package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Repository layer

// Task represents a to-do task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//Description string `json:"description"`
	Status bool `json:"status"`
}

// TaskRepository is a repository interface for tasks
type TaskRepository interface {
	GetTasks() ([]Task, error)
	GetTask(id int) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(task Task) (Task, error)
	DeleteTask(id int) error
}

// FileTaskRepository is a file-based implementation of TaskRepository
type FileTaskRepository struct {
	FilePath string
}

// GetTasks returns all tasks from the repository
func (repo *FileTaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.Open(repo.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(content, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *FileTaskRepository) SaveTasks(tasks []Task) error {
	file, err := os.Create(repo.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(repo.FilePath, data, 0666)
	if err != nil {
		panic(err)
	}

	return nil
}

// GetTask returns a single task by its ID
func (repo *FileTaskRepository) GetTask(id int) (Task, error) {
	var task Task

	tasks, err := repo.GetTasks()
	if err != nil {
		return task, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return task, nil
}

// CreateTask adds a new task to the repository
func (repo *FileTaskRepository) CreateTask(task Task) (Task, error) {
	tasks, err := repo.GetTasks()
	if err != nil {
		return task, err
	}

	task.ID = len(tasks) + 1
	tasks = append(tasks, task)

	if err := repo.SaveTasks(tasks); err != nil {
		return task, err
	}

	return task, nil
}

// UpdateTask updates an existing task in the repository
func (repo *FileTaskRepository) UpdateTask(task Task) (Task, error) {
	tasks, err := repo.GetTasks()
	if err != nil {
		return task, err
	}

	found := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			found = true
			break
		}
	}
	if !found {
		return task, fmt.Errorf("task not found")
	}

	if err := repo.SaveTasks(tasks); err != nil {
		return task, err
	}

	return task, nil
}

func (repo *FileTaskRepository) DeleteTask(id int) error {
	tasks, err := repo.GetTasks()
	found := false
	if err != nil {
		return err
	}

	newTasks := make([]Task, 0)

	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("task not found")
	}

	if err := repo.SaveTasks(newTasks); err != nil {
		return err
	}

	return nil
}
