package store

import (
	"fmt"
	"sync"
)

type OrderStorager interface {
	Get(OrderID int) (Order, error)
	Create(Order Order) Order
	Delete(OrderID int) error
	GetList() []Order
}

type OrderStorage struct {
	primaryKeyIDx map[int]*Order
	autoIncrement int
	sync.Mutex
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{
		primaryKeyIDx: make(map[int]*Order, 13),
	}
}

func (p *OrderStorage) Create(Order Order) Order {
	p.Lock()
	defer p.Unlock()
	Order.ID = p.autoIncrement
	p.primaryKeyIDx[Order.ID] = &Order
	p.autoIncrement++

	return Order
}

func (p *OrderStorage) Get(OrderID int) (Order, error) {
	if v, ok := p.primaryKeyIDx[OrderID]; ok {
		return *v, nil
	}
	return Order{}, fmt.Errorf("not found")
}

func (p *OrderStorage) Delete(OrderID int) error {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.primaryKeyIDx[OrderID]; ok {
		delete(p.primaryKeyIDx, OrderID)
		return nil
	}

	return fmt.Errorf("not found")
}
func (p *OrderStorage) GetList() []Order {
	data := []Order{}

	for _, v := range p.primaryKeyIDx {
		data = append(data, *v)
	}
	return data
}
