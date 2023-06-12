package pet

import (
	"fmt"
	"sync"
)

type PetStorager interface {
	Create(pet Pet) Pet
	Update(name, status string, petID int) error
	FullUpdate(pet Pet) (Pet, error)
	Delete(petID int) error
	GetByID(petID int) (Pet, error)
	GetList() []Pet
	AddImage(petID int, img string) error
	FindByStatus(status []string) []Pet
}

type PetStorage struct {
	data          []*Pet
	primaryKeyIDx map[int]*Pet
	autoIncrement int
	sync.Mutex
}

func NewPetStorage() *PetStorage {
	return &PetStorage{
		data:          make([]*Pet, 0, 13),
		primaryKeyIDx: make(map[int]*Pet, 13),
	}
}

func (p *PetStorage) Create(pet Pet) Pet {
	p.Lock()
	defer p.Unlock()
	pet.ID = p.autoIncrement
	p.primaryKeyIDx[pet.ID] = &pet
	p.autoIncrement++
	p.data = append(p.data, &pet)

	return pet
}

func (p *PetStorage) Update(name, status string, petID int) error {
	v, ok := p.primaryKeyIDx[petID]
	if !ok {
		return fmt.Errorf("not found")
	}

	v.Name = name
	v.Status = status
	return nil
}

func (p *PetStorage) FullUpdate(pet Pet) (Pet, error) {
	_, ok := p.primaryKeyIDx[pet.ID]
	if !ok {
		return Pet{}, fmt.Errorf("not found")
	}

	p.primaryKeyIDx[pet.ID] = &pet
	for k, v := range p.data {
		if v.ID == pet.ID {
			p.data[k] = &pet
			break
		}
	}
	return pet, nil
}

func (p *PetStorage) Delete(petID int) error {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.primaryKeyIDx[petID]; !ok {
		return fmt.Errorf("not found")
	}
	delete(p.primaryKeyIDx, petID)
	return nil
}

func (p *PetStorage) GetByID(petID int) (Pet, error) {
	if v, ok := p.primaryKeyIDx[petID]; ok {
		return *v, nil
	}
	return Pet{}, fmt.Errorf("not found")
}

func (p *PetStorage) GetList() []Pet {
	pets := make([]Pet, 0)

	for _, v := range p.data {
		pets = append(pets, *v)
	}
	return pets
}

func (p *PetStorage) AddImage(petID int, img string) error {
	p.Lock()
	defer p.Unlock()
	v, ok := p.primaryKeyIDx[petID]
	if !ok {
		return fmt.Errorf("not found")
	}
	images := v.PhotoUrls
	images = append(images, img)
	v.PhotoUrls = images

	return nil
}

func (p *PetStorage) FindByStatus(filters []string) []Pet {
	pets := p.GetList()
	ans := make([]Pet, 0)
	for _, v := range pets {
		for _, status := range filters {
			if v.Status == status {
				ans = append(ans, v)
			}
		}
	}
	return ans
}
