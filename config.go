package config

import (
	"errors"
	"os"
)

type Provider interface {
	SetRawString(raw string)
	SetRawBytes(raw []byte)
	Get(data interface{}) error
}

var (
	Prefix string
)

func New(file, format string) (provider Provider, err error) {

	// set config file full path
	filepath := Prefix + file
	fp, fperr := os.Open(filepath)
	if fperr != nil {
		return nil, fperr
	}

	// read config file raw data.
	fstat, _ := fp.Stat()
	raw := make([]byte, fstat.Size())
	fp.Read(raw)

	switch format {
	case "json":
		provider = &JSON{}
		provider.SetRawBytes(raw)

	case "xml":
		provider = &XML{}
		provider.SetRawBytes(raw)

	default:
		return nil, errors.New("Not support this format config file.")
	}
	return
}
