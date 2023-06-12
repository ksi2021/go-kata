package user

import (
	"fmt"
	"sync"
)

type UserStorager interface {
	Create(User User) User
	Update(username string, user User) (User, error)
	Delete(username string) error
	GetByUsername(Username string) (User, error)
	CreateWithList(users []User) error
}

type UserStorage struct {
	primaryKeyIDx map[string]*User
	autoIncrement int
	sync.Mutex
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		primaryKeyIDx: make(map[string]*User, 13),
	}
}

func (p *UserStorage) Create(User User) User {
	p.Lock()
	defer p.Unlock()
	User.ID = p.autoIncrement
	p.primaryKeyIDx[User.Username] = &User
	p.autoIncrement++

	return User
}

func (p *UserStorage) CreateWithList(users []User) error {
	p.Lock()
	defer p.Unlock()

	for _, v := range users {
		v.ID = p.autoIncrement
		p.primaryKeyIDx[v.Username] = &v
		p.autoIncrement++
	}

	return nil
}

func (p *UserStorage) Update(username string, user User) (User, error) {
	_, ok := p.primaryKeyIDx[username]
	if !ok {
		return User{}, fmt.Errorf("not found")
	}

	p.Lock()
	defer p.Unlock()

	delete(p.primaryKeyIDx, username)
	user.ID = p.autoIncrement
	p.autoIncrement++
	p.primaryKeyIDx[user.Username] = &user

	return user, nil
}

func (p *UserStorage) Delete(username string) error {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.primaryKeyIDx[username]; !ok {
		return fmt.Errorf("not found")
	}
	delete(p.primaryKeyIDx, username)
	return nil
}

func (p *UserStorage) GetByUsername(Username string) (User, error) {
	if v, ok := p.primaryKeyIDx[Username]; ok {
		return *v, nil
	}
	return User{}, fmt.Errorf("not found")
}
