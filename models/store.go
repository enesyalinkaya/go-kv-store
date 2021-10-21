package models

import memory "github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"

type storeModel struct {
	db *memory.MemoryClient
}

// StoreModel interface represents store model functionality.
type StoreModel interface {
	Set(key string, value string) string
	Get(key string) string
	Flush()
}

func NewStoreModel(db *memory.MemoryClient) StoreModel {
	return &storeModel{db: db}
}

func (m *storeModel) Set(key string, value string) string {
	m.db.Set(key, value)
	return value
}

func (m *storeModel) Get(key string) string {
	return m.db.Get(key)
}

func (m *storeModel) Flush() {
	m.db.Flush()
}
