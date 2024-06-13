package gocuddle

import "time"

type Session struct {
	ID string
	Values map[string]interface{}
	Expires time.Time
	New bool
	Changed bool
}
