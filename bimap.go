package bimap

import "sync"

type biMap struct {
	s sync.RWMutex
	forward map[interface{}]interface{}
	inverse map[interface{}]interface{}
}

func NewBiMap() *biMap {
	return &biMap{forward: make(map[interface{}]interface{}), inverse: make(map[interface{}]interface{})}
}


func (b *biMap) Insert(k interface{}, v interface{}) {
	b.s.Lock()
	defer b.s.Unlock()
	b.forward[k] = v
	b.inverse[v] = k
}

func (b *biMap) Exists(k interface{}) bool {
	b.s.RLock()
	defer b.s.RUnlock()
	_, ok := b.forward[k]
	return ok
}

func (b *biMap) InverseExists(k interface{}) bool {
	b.s.RLock()
	defer b.s.RUnlock()

	_, ok := b.inverse[k]
	return ok
}

func (b *biMap) Get(k interface{}) (interface{}, bool) {
	if b.Exists(k) {
		b.s.RLock()
		defer b.s.RUnlock()
		return b.forward[k], true
	}
	return "", false
}

func (b *biMap) InverseGet(v interface{}) (interface{}, bool) {
	if b.InverseExists(v) {
		b.s.RLock()
		defer b.s.RUnlock()
		return b.inverse[v], true
	}
	return "", false
}

func (b *biMap) Delete(k interface{}) {
	if b.Exists(k) {
		val, _ := b.Get(k)
		b.s.Lock()
		defer b.s.Unlock()
		delete(b.forward, k)
		delete(b.inverse, val)
	}
}

func (b *biMap) InverseDelete(v interface{}) {
	if b.InverseExists(v) {
		key, _ := b.InverseGet(v)
		b.s.Lock()
		defer b.s.Unlock()
		delete(b.inverse, v)
		delete(b.forward, key)
	}
}

func (b*biMap) Size() int{
	b.s.RLock()
	defer b.s.RUnlock()
	return len(b.forward)
}