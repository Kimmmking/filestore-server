package store

import (
	"filestore-server/store"
	"filestore-server/store/factory"
	"sync"
)

type MemStore struct {
	sync.RWMutex
	filemetas map[string]*store.FileMeta
}

func init() {
	factory.Register("mem", &MemStore{
		filemetas: make(map[string]*store.FileMeta),
	})
}

func (ms *MemStore) Create(filemeta *store.FileMeta) error {
	return nil
}

func (ms *MemStore) Update(filemeta *store.FileMeta) error {
	return nil
}

func (ms *MemStore) Get(string) (store.FileMeta, error) {
	return store.FileMeta{}, nil
}

func (ms *MemStore) GetAll() ([]store.FileMeta, error) {
	return []store.FileMeta{}, nil
}

func (ms *MemStore) Delete(name string) error {
	return nil
}
