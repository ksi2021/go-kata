package repository

import (
	"encoding/json"
	"io"
)

type Repository interface {
	Save(record interface{}) error
	Find(id int) (interface{}, error)
	FindAll() ([]interface{}, error)
}

type UserRepository struct {
	File io.ReadWriter
}

func NewUserRepository(file io.ReadWriter) *UserRepository {
	return &UserRepository{
		File: file,
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Загружать данные в конструкторе

func (r *UserRepository) Save(record interface{}) error {
	data := record.(User)
	converBytes, er := json.Marshal(data)
	if er != nil {
		panic(er)
	}
	_, err := r.File.Write(converBytes)
	if err != nil {
		panic(err)
	}

	_, _ = r.File.Write([]byte(","))

	return err
}

func (r *UserRepository) Find(id int) (interface{}, bool) {
	data, err := r.FindAll()

	if !err {
		return nil, false
	}

	for _, v := range data {
		if v.(User).ID == id {
			return v, true
		}
	}

	return nil, false
}

func (r *UserRepository) FindAll() ([]interface{}, bool) {
	var records []User

	data, err := io.ReadAll(r.File)
	if err != nil && err != io.EOF {
		return nil, false
	}

	newData := make([]byte, len(data)+1)
	newData[0] = 91
	copy(newData[1:], data)
	newData[len(newData)-1] = 93

	err = json.Unmarshal(newData, &records)
	if err != nil {
		return nil, false
	}

	var interfaceSlice []interface{} = make([]interface{}, len(records))
	for i, d := range records {
		interfaceSlice[i] = d
	}
	return interfaceSlice, true
}
