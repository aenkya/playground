package model

import (
	"errors"
	"fmt"
)

type ToDo struct {
	Label string `json:"label"`
	ID    int    `json:"id"`
}

func (t *ToDo) String() string {
	return fmt.Sprintf("ID: %d, Label: %v", t.ID, t.Label)
}

func NewToDo(label string) (*ToDo, error) {
	if label == "" {
		return nil, errors.New("label can't be empty")
	}

	return &ToDo{
		Label: label,
	}, nil
}
