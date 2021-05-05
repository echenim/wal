package server

import "sync"

type Log struct {
	mu      sync.Mutex
	records []Record
}
