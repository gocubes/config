package config

import (
	"encoding/json"
)

// parse json string
type JSON struct {
	raw []byte
}

func (j *JSON) SetRawString(raw string) {
	j.raw = []byte(raw)
}

func (j *JSON) SetRawBytes(raw []byte) {
	j.raw = raw
}

func (j *JSON) Get(data interface{}) error {
	return json.Unmarshal(j.raw, data)
}
