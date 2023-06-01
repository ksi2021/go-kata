package repository

import (
	. "github.com/ksi2021/go-kata/module3/clean_architecture/service/repo"
	. "github.com/ksi2021/go-kata/module3/clean_architecture/service/service"
	"testing"
)

func TestGetList(t *testing.T) {

	tests := []struct {
		name string
		want error
	}{
		{
			name: "GetList 1",
			want: nil,
		},
		{
			name: "GetList 2",
			want: nil,
		},
		{
			name: "GetList 3",
			want: nil,
		},
	}

	service := MakeService("test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, err := service.ListTodos(); err != tt.want {
				t.Errorf("Data = %v, result = %v , want %v", data, err, tt.want)
			}
		})
	}
}

func TestCreateTodo(t *testing.T) {

	tests := []struct {
		name string
		data string
		want error
	}{
		{
			name: "CreateTodo 1",
			data: "data #1",
			want: nil,
		},
		{
			name: "CreateTodo 2",
			data: "data #2",
			want: nil,
		},
		{
			name: "CreateTodo 3",
			data: "data #3",
			want: nil,
		},
	}

	service := MakeService("test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.CreateTodo(tt.data); err != tt.want {
				t.Errorf("result = %v , want %v", err, tt.want)
			}
		})
	}
}

func TestCompleteTodo(t *testing.T) {

	tests := []struct {
		name string
		data Task
		want error
	}{
		{
			name: "CompleteTodo 2 #1",
			data: Task{Title: "complete 1", ID: 1},
			want: nil,
		},
		{
			name: "CompleteTodo 2",
			data: Task{Title: "complete 2", ID: 2},
			want: nil,
		},
		{
			name: "CompleteTodo 2",
			data: Task{Title: "complete 2", ID: 3},
			want: nil,
		},
	}

	service := MakeService("test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.CompleteTodo(tt.data); err != tt.want {
				t.Errorf("result = %v , want %v", err, tt.want)
			}
		})
	}
}
func TestRemoveTodo(t *testing.T) {

	tests := []struct {
		name string
		data Task
		want error
	}{
		{
			name: "RemoveTodo 1",
			data: Task{Title: "complete 3", ID: 3},
			want: nil,
		},
		{
			name: "RemoveTodo 2",
			data: Task{Title: "complete 3", ID: 100},
			want: nil,
		},
	}

	service := MakeService("test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.RemoveTodo(tt.data); err != tt.want {
				t.Errorf("result = %v , want %v", err, tt.want)
			}
		})
	}
}
