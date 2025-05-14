package db

import "enkya.org/playground/cmd/todo/model"

type IMemory interface {
	Add(*model.ToDo)
	Update(*model.ToDo)
	Delete(*model.ToDo)
	GetByID(id int) *model.ToDo
	GetAll() []*model.ToDo
}

type memory struct {
	mem     map[int]*model.ToDo
	counter int
}

func New() IMemory {
	return &memory{
		mem:     make(map[int]*model.ToDo),
		counter: 0,
	}
}

func (m *memory) Add(t *model.ToDo) {
	m.counter++
	t.ID = m.counter
	m.mem[t.ID] = t
}

func (m *memory) Update(t *model.ToDo) {
	m.mem[t.ID] = t
}

func (m *memory) Delete(t *model.ToDo) {
	delete(m.mem, t.ID)
}

func (m *memory) GetByID(id int) *model.ToDo {
	val, ok := m.mem[id]
	if !ok {
		return nil
	}

	return val
}

func (m *memory) GetAll() []*model.ToDo {
	todos := make([]*model.ToDo, 0)
	for _, v := range m.mem {
		todos = append(todos, v)
	}

	return todos
}
