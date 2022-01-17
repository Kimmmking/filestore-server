package factory

import (
	"filestore-server/store"
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.FileMetaStore)
)

func Register(name string, p store.FileMetaStore) {
	providersMu.Lock()
	defer providersMu.Unlock()

	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = p
}

func New(providerName string) (store.FileMetaStore, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}

	return p, nil
}
