package config

import (
	"encoding/json"
	"os"
)

// parse json string
type JSON struct {
	path string
	raw  []byte
}

func (j *JSON) SetRawString(raw string) {
	j.raw = []byte(raw)
}

func (j *JSON) SetRawBytes(raw []byte) {
	j.raw = raw
}

func (j *JSON) GetRawBytes() []byte {
	return j.raw
}

func (j *JSON) Get(data interface{}) error {
	return json.Unmarshal(j.raw, data)
}

func (j *JSON) Set(data interface{}) (int, error) {
	var err error
	// open file
	fp, fperr := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer fp.Close()
	if fperr != nil {
		return 0, fperr
	}

	// encode
	if j.raw, err = json.MarshalIndent(&data, "", "    "); err != nil {
		return 0, err
	}
	j.raw = append(j.raw, '\n')

	// write
	return fp.Write(j.raw)
}

func (j *JSON) SetPath(path string) {
	j.path = path
}

func (j *JSON) GetPath() string {
	return j.path
}

func (j *JSON) Reload(data interface{}) error {
	fp, fperr := os.OpenFile(j.GetPath(), os.O_RDONLY|os.O_CREATE, os.ModePerm)

	if fperr != nil {
		return fperr
	}

	// read config file raw data.
	fstat, _ := fp.Stat()
	raw := make([]byte, fstat.Size())
	fp.Read(raw)

	j.SetRawBytes(raw)
	return j.Get(data)
}
