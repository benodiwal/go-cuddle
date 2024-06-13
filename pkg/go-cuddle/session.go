package gocuddle

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

type Session struct {
	ID string
	Values map[string]interface{}
	Expires time.Time
	New bool
	Changed bool
}

func encodeSession(session *Session) string {
	data, err := json.Marshal(session.Values)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(data)
}


func decodeSession(encoded string) (map[string]interface{}, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var values map[string]interface{}
	err = json.Unmarshal(data, &values)
	return values, err
}
