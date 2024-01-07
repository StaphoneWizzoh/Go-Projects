package recipes

import "errors"

var (
	NotFoundError = errors.New("Not Found")
)

type MemStore struct {
	list map[string]Recipe
}

func NewMemStore() *MemStore {
	list := make(map[string]Recipe)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, recipe Recipe) error {
	m.list[name] = recipe
	return nil
}

func (m MemStore) Get(name string) (Recipe, error) {
	if val, ok := m.list[name]; ok {
		return val, nil
	}
	return Recipe{}, NotFoundError
}

func (m MemStore) List() (map[string]Recipe, error) {
	return m.list, nil
}

func (m MemStore) Update(name string, recipe Recipe) error {
	if _, ok := m.list[name]; ok {
		m.list[name] = recipe
		return nil
	}
	return NotFoundError
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
