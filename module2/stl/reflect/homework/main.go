package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var reflectedStructs map[string][]Field

type Field struct {
	Name string
	Tags map[string]string `fakesize:"2,5"`
}

type IUser struct {
	UserDTO
	EmailVerifyDTO
}

type IUserArr []IUser

type UserDTO struct {
	ID            int       `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	Name          string    `json:"name" db:"name" db_type:"varchar(55)" db_default:"default null" db_ops:"create,update"`
	Phone         string    `json:"phone" db:"phone" db_type:"varchar(34)" db_default:"default null" db_index:"index,unique" db_ops:"create,update"`
	Email         string    `json:"email" db:"email" db_type:"varchar(89)" db_default:"default null" db_index:"index,unique" db_ops:"create,update"`
	Password      string    `json:"password" db:"password" db_type:"varchar(144)" db_default:"default null" db_ops:"create,update"`
	Status        int       `json:"status" db:"status" db_type:"int" db_default:"default 0" db_ops:"create,update"`
	Role          int       `json:"role" db:"role" db_type:"int" db_default:"not null" db_ops:"create,update"`
	Verified      bool      `json:"verified" db:"verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	EmailVerified bool      `json:"email_verified" db:"email_verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	PhoneVerified bool      `json:"phone_verified" db:"phone_verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	CreatedAt     time.Time `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index"`
}

type EmailVerifyDTO struct {
	ID        int       `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	Email     string    `json:"email" db:"email" db_type:"varchar(89)" db_default:"not null" db_index:"index,unique" db_ops:"create,update"`
	UserID    int       `json:"user_id,omitempty" db:"user_id" db_ops:"create" db_type:"int" db_default:"not null" db_index:"index"`
	Hash      string    `json:"hash,omitempty" db:"hash" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index"`
	Verified  bool      `json:"verified" db:"verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	CreatedAt time.Time `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	structs := []interface{}{
		UserDTO{},
		EmailVerifyDTO{},
	}

	r := reflect.ValueOf(structs)
	fmt.Printf("%+v\n\n", r)

	_map := make(map[string]IUserArr)
	r_map := reflect.ValueOf(_map)
	r_map.SetMapIndex(reflect.ValueOf("test"), reflect.ValueOf(IUserArr{IUser{}}))
	r_map.SetMapIndex(reflect.ValueOf("test2"), reflect.ValueOf(IUserArr{IUser{}}))
	fmt.Printf("%+v\n\n", r_map)

	// заполни данными структур согласно заданиям
	reflectedStructs = make(map[string][]Field, len(structs))
	for i := 0; i < len(structs); i++ {
		n := "struct" + fmt.Sprint(i+1)
		f := []Field{}
		f_ := Field{}
		for i := 0; i < rand.Intn(9)+1; i++ {
			gofakeit.Struct(&f_)
			f = append(f, f_)
		}
		reflectedStructs[n] = f

	}

	fmt.Printf("%+v \n", reflectedStructs)
}
