package cachemanager

import "errors"

// Vars
var (
	ErrInvalidHandler = errors.New("Invalid handler")
	ErrCacheMiss = errors.New("Cache miss")
)

// Manager represents a cache manager capable of switching between several cache handlers
type Manager interface {
	AddHandler(n string, h Handler) Manager
	Del(k string) error
	Get(k string) ([]byte, error)
	GetHandler(n string) (Handler, error)
}

// NewManager creates a new cache manager
func NewManager() Manager {
	return &manager{
		handlers: make(map[string]Handler),
	}
}

type manager struct {
	handlers map[string]Handler
}

// AddHandler adds a new handler
func (m *manager) AddHandler(n string, h Handler) Manager {
	m.handlers[n] = h
	return m
}

func (m manager) Del(k string) error {
	var e error
	for _, h := range m.handlers {
		e = h.Del(k)
		if e != nil {
			return e
		}
	}
	return e
}

func (m manager) Get(k string) ([]byte, error) {
	var o []byte
	var e error
	for _, h := range m.handlers {
		o, e = h.Get(k)
		if e == nil || (e != nil && e.Error() != ErrCacheMiss.Error()) {
			return o, e
		}
	}
	return o, e
}

// GetHandler returns the handler
func (m manager) GetHandler(n string) (Handler, error) {
	var e error
	if h, ok := m.handlers[n]; ok {
		return h, e
	}
	return nil, ErrInvalidHandler
}
