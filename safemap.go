package safemap

import (
	"iter"
	"maps"
	"sync"
)

type Map[K comparable, V any] struct {
	m    map[K]V
	rwmu *sync.RWMutex
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:    map[K]V{},
		rwmu: &sync.RWMutex{},
	}
}

func (m *Map[K, V]) All() iter.Seq2[K, V] {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()

	return maps.All(m.m)
}

func (m *Map[K, V]) Delete(k K) {
	m.rwmu.Lock()
	defer m.rwmu.Unlock()

	delete(m.m, k)
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()

	v, ok := m.m[k]
	return v, ok
}

func (m *Map[K, V]) Insert(k K, v V) {
	m.rwmu.Lock()
	defer m.rwmu.Unlock()

	m.m[k] = v
}
