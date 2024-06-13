package gocuddle

import (
	"net/http"
	"time"
)

type Option func(*Manager)

func WithName(name string) Option {
	return func(m *Manager) {
		m.name = name
	}
}

func WithKeys(keys []string) Option {
	return func(m *Manager) {
		m.keys = keys
	}
}

func WithSecure(secure bool) Option {
	return func(m *Manager) {
		m.secure = secure
	}
}

func WithHTTPOnly(httpOnly bool) Option {
    return func(m *Manager) {
        m.httpOnly = httpOnly
    }
}

func WithSameSite(sameSite http.SameSite) Option {
    return func(m *Manager) {
        m.sameSite = sameSite
    }
}

func WithMaxAge(maxAge time.Duration) Option {
    return func(m *Manager) {
        m.lifetime = maxAge
    }
}
