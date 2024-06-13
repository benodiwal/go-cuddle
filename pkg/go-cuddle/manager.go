package gocuddle

import (
	"net/http"
	"time"
)

type Manager struct {
	sessions map[string]*Session
	lifetime time.Duration
	keys []string
	name string
	secure bool
	httpOnly bool
	sameSite http.SameSite
}

func NewManager(options ...Option) *Manager {
	m := &Manager{
		sessions: make(map[string]*Session),
		lifetime: 24 * time.Hour,
	}

	for _, option := range options {
		option(m)
	}

	return m
}
