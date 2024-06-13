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
