package cache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
)

var errFileNotCached = errors.New("File not found in cache store")
var once sync.Once
var managerInstance Manager

// Resource is used to store a cached resource
type Resource struct {
	content []byte
}

// Manager can be used to manage cached resources
type Manager struct {
	store map[string]Resource
}

// CacheFile is used to cache a given file with a given path relative to where the binary was executed
func (m Manager) CacheFile(path string, id string) (bool, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}
	m.store[id] = Resource{content: content}
	return true, nil
}

// GetFile is used to get a cached file from cache store
func (m Manager) GetFile(id string) ([]byte, error) {
	c, ok := m.store[id]
	if !ok {
		return nil, fmt.Errorf("File not found in cache store: %v: %w", id, errFileNotCached)
	}
	return c.content, nil
}

// GetManager returns a single instance of the manager used in the currently running process
func GetManager() Manager {
	once.Do(func() {
		managerInstance = Manager{store: make(map[string]Resource)}
	})
	return managerInstance
}
