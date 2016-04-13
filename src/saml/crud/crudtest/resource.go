package crudtest

import (
	"saml/crud"
)

// InMemoryResource ...
type InMemoryResource struct {
	Key   string
	Value string
}

// GetID ...
func (r *InMemoryResource) GetID() string {
	return r.Key
}

// SetID ...
func (r *InMemoryResource) SetID(id string) {
	r.Key = id
}

// InMemoryResourceManager ...
type InMemoryResourceManager struct {
	Map    map[string]*InMemoryResource
	NextID int
}

// Get ...
func (m *InMemoryResourceManager) Get(id string) (crud.Resource, error) {
	r, ok := m.Map[id]
	if !ok {
		return nil, crud.ErrNotFound
	}
	return r, nil
}

// Post ...
func (m *InMemoryResourceManager) Post(r *InMemoryResource) (crud.Resource, error) {
	r.SetID(string(m.NextID))
	m.Map[r.GetID()] = r
	m.NextID++
	return r, nil
}

// Put ...
func (m *InMemoryResourceManager) Put(r *InMemoryResource) (crud.Resource, error) {
	m.Map[r.GetID()] = r
	return r, nil
}

// Delete ...
func (m *InMemoryResourceManager) Delete(r crud.Resource) error {
	delete(m.Map, r.GetID())
	return nil
}
