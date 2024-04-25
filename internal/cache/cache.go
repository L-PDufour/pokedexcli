package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data  map[string]CacheEntry
	mutex *sync.Mutex
}

func NewCache(expirationInterval time.Duration) Cache {
	m := Cache{
		data:  make(map[string]CacheEntry),
		mutex: &sync.Mutex{},
	}
	go m.reapLoop(expirationInterval)
	return m
}

func (m *Cache) Add(key string, val []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (m *Cache) Get(key string) ([]byte, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	entry, ok := m.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (m *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		m.reap(time.Now().UTC(), interval)
	}
}

func (m *Cache) reap(now time.Time, last time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(m.data, k)
		}
	}
}
