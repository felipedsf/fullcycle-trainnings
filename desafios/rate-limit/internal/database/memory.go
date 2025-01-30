package database

import (
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/dto"
	"sync"
)

var once sync.Once

type Memory struct {
	database map[string]*dto.RateLimit
}

func GetMemoryDb() *Memory {
	var instance Memory
	once.Do(func() {
		instance = Memory{
			database: make(map[string]*dto.RateLimit),
		}
	})
	return &instance
}

func (m *Memory) Get(id string) *dto.RateLimit {
	if val, ok := m.database[id]; ok {
		return val
	}
	return nil
}

func (m *Memory) Save(id string, limit *dto.RateLimit) {
	m.database[id] = limit
}
