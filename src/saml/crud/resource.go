package crud

import (
	"fmt"
)

var (
	// ErrNotFound is error returned when resource isn't found.
	ErrNotFound = fmt.Errorf("Not Found")
)

// ResourceManager ...
type ResourceManager interface {
	Get(id string) (Resource, error)
	Post(resource Resource) (Resource, error)
	Put(resource Resource) (Resource, error)
	Delete(resource Resource) error
}

// Resource ...
type Resource interface {
	GetID() string
	SetID(id string)
}
