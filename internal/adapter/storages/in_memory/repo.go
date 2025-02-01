package in_memory

import (
	"sync"
)

type Repo struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewRepo() *Repo {
	return &Repo{
		data: make(map[string]string)}
}
