package config

import (
	"encoding/xml"
)

// parse xml string
type XML struct {
	raw []byte
}

func (x *XML) SetRawString(raw string) {
	x.raw = []byte(raw)
}

func (x *XML) SetRawBytes(raw []byte) {
	x.raw = raw
}

func (x *XML) Get(data interface{}) error {
	return xml.Unmarshal(x.raw, data)
}
