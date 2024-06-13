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
		name: "cookie-session",
		sessions: make(map[string]*Session),
		lifetime: 24 * time.Hour,
	}

	for _, option := range options {
		option(m)
	}

	return m
}

func (m *Manager) NewSession(w http.ResponseWriter) *Session {
	id := generateSessionID()
	session := &Session {
		ID: id,
		Values: make(map[string]interface{}),
		Expires: time.Now().Add(m.lifetime),
		New: true,
		Changed: true,
	}

	m.sessions[id] = session

	http.SetCookie(w, &http.Cookie{
		Name: m.name,
		Value: id,
		Expires: session.Expires,
		Path: "/",
		Secure: m.secure,
		HttpOnly: m.httpOnly,
		SameSite: m.sameSite,
	})

	return session
}

func (m *Manager) GetSession(r *http.Request) *Session {
	cookie, err := r.Cookie(m.name)
	if err != nil {
		return nil
	}
	session, exists := m.sessions[cookie.Value]
	if !exists || session.Expires.Before(time.Now()) {
		return nil
	}
	return session
}

func (m *Manager) SaveSession(w http.ResponseWriter, session *Session) {
    session.Expires = time.Now().Add(m.lifetime)
    http.SetCookie(w, &http.Cookie{
        Name:     m.name,
        Value:    session.ID,
        Expires:  session.Expires,
        Path:     "/",
        Secure:   m.secure,
        HttpOnly: m.httpOnly,
        SameSite: m.sameSite,
    })
    m.sessions[session.ID] = session
}

func (m *Manager) DestroySession(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie(m.name)
    if err != nil {
        return
    }

    delete(m.sessions, cookie.Value)
    http.SetCookie(w, &http.Cookie{
        Name:   m.name,
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    })
}
