package singleton

import "sync"

var once sync.Once
var mutex sync.Mutex
var instance *singleton

type singleton struct {
	id int
}

func GetSingletonInstance() *singleton {
	once.Do(func() {
		instance = &singleton{id: 0}
	})
	return instance
}

func (s *singleton) GetNextId() int {
	defer mutex.Unlock()
	mutex.Lock()
	s.id = s.id + 1
	return s.id
}
