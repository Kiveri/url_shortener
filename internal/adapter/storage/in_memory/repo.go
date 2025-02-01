package in_memory

import (
	"sync"
)

type Repo struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewRepo() *Repo {
	return &Repo{
		urls: make(map[string]string),
	}
}
