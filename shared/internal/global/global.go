package global

import (
	"sync"
)

type set struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func (s *set) Get(name string) (value interface{}) {
	s.mu.RLock()
	value = s.data[name]
	s.mu.RUnlock()
	return
}

func (s *set) Set(name string, value interface{}) {
	s.mu.Lock()
	s.data[name] = value
	s.mu.Unlock()
}

func (s *set) GetOrSet(name string, fn func() interface{}) interface{} {
	s.mu.RLock()
	value, ok := s.data[name]
	s.mu.RUnlock()

	if ok {
		return value
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if value = fn(); value != nil {
		s.data[name] = value
	}

	return value
}

var s = &set{data: make(map[string]interface{})}

func Get(name string) interface{} {
	return s.Get(name)
}

func Set(name string, value interface{}) {
	s.Set(name, value)
}

func GetOrSet(name string, fn func() interface{}) interface{} {
	return s.GetOrSet(name, fn)
}
