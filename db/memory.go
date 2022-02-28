package db

import (
	"fmt"

	"github.com/JosacabDev/go_api/model"
)

type Memory struct {
	ID     int
	Peopel map[int]model.Person
}

func NewMemory() Memory {
	people := make(map[int]model.Person)

	return Memory{
		ID:     0,
		Peopel: people,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonNotNil
	}

	m.ID++
	m.Peopel[m.ID] = *person
	return nil
}

func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonNotNil
	}

	if _, ok := m.Peopel[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDDoestExist)
	}

	m.Peopel[ID] = *person
	return nil
}

func (m *Memory) Delete(ID int) error {
	if _, ok := m.Peopel[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDDoestExist)
	}

	delete(m.Peopel, ID)
	return nil
}

func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Peopel[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDDoestExist)
	}

	return person, nil
}

func (m *Memory) GetAll() (model.Peopel, error) {
	var result model.Peopel
	for _, v := range m.Peopel {
		result = append(result, v)
	}

	return result, nil
}
