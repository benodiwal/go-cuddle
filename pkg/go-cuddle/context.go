package gocuddle

import (
	"context"
	"net/http"
)

func WithSession(r *http.Request, session *Session) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), "session", session))
}

func GetSession(r *http.Request) *Session {
	return r.Context().Value("session").(*Session)
}
