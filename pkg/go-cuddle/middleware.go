package gocuddle

import "net/http"

func (m *Manager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := m.GetSession(r)
		if session == nil {
			session = m.NewSession(w)
		}

		r = WithSession(r, session)
		next.ServeHTTP(w, r)

		if session.Changed {
			m.SaveSession(w, session)
		}
	})
}
