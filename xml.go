package config

import (
	"encoding/xml"
	"os"
)

// parse xml string
type XML struct {
	path string
	raw  []byte
}

func (x *XML) SetRawString(raw string) {
	x.raw = []byte(raw)
}

func (x *XML) SetRawBytes(raw []byte) {
	x.raw = raw
}

func (x *XML) GetRawBytes() []byte {
	return x.raw
}

func (x *XML) Get(data interface{}) error {
	return xml.Unmarshal(x.raw, data)
}

func (x *XML) Set(data interface{}) (int, error) {
	// open file
	fp, fperr := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer fp.Close()
	if fperr != nil {
		return 0, fperr
	}

	// encode
	bytes, err := xml.MarshalIndent(&data, "", "    ")
	if err != nil {
		return 0, err
	}
	bytes = append(bytes, '\n')

	// add xml header
	x.raw = make([]byte, len(xml.Header)+len(bytes))
	copy(x.raw, []byte(xml.Header))
	copy(x.raw[len(xml.Header):], bytes)

	// write
	return fp.Write(x.raw)
}

func (x *XML) SetPath(path string) {
	x.path = path
}
