package server

import "sync"

type Log struct {
	Mu      sync.Mutex
	Records []Record
}
