package service

import (
	"errors"

	"enkya.org/playground/cmd/todo/db"
	"enkya.org/playground/cmd/todo/model"
)

type IToDoService interface {
	Add(*model.ToDo) (*model.ToDo, error)
	Update(*model.ToDo) (*model.ToDo, error)
	Delete(id int) error
	GetByID(id int) *model.ToDo
	GetAll() []*model.ToDo
}

type toDoService struct {
	db db.IMemory
}

func NewService(db db.IMemory) IToDoService {
	return &toDoService{
		db: db,
	}
}

func (s *toDoService) Add(t *model.ToDo) (*model.ToDo, error) {
	s.db.Add(t)
	return t, nil
}

func (s *toDoService) Update(t *model.ToDo) (*model.ToDo, error) {
	found := s.db.GetByID(t.ID)
	if found != nil {
		found.Label = t.Label
		s.db.Update(found)

		return found, nil
	}

	return nil, errors.New("todo not found")
}

func (s *toDoService) Delete(id int) error {
	found := s.db.GetByID(id)
	if found == nil {
		return errors.New("todo not found")
	}

	s.db.Delete(found)

	return nil
}

func (s *toDoService) GetByID(id int) *model.ToDo {
	return s.db.GetByID(id)
}

func (s *toDoService) GetAll() []*model.ToDo {
	return s.db.GetAll()
}
