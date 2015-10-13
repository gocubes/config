package config

import (
	"errors"
	"fmt"
	"os"
)

type Provider interface {
	SetRawString(raw string)
	SetRawBytes(raw []byte)
	Get(data interface{}) error
	Set(data interface{}) (int, error)
}

var (
	Prefix   string
	filepath string
	provider Provider
)

func New(file, format string) (provider Provider, err error) {

	// set config file full path
	filepath = Prefix + file
	fp, fperr := os.Open(filepath)
	if os.IsNotExist(fperr) {
		fperr = nil
		filepath = Prefix + "/etc/" + file
		fp, fperr = os.Open(filepath)
	}

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

// init default provider
func defaultProvider() {
	var derr error
	Prefix = "./"
	provider, derr = New("main.json", "json")
	if derr != nil {
		fmt.Printf("[config package]get default config error,details: %v\n", derr)
		os.Exit(1)
	}
}

func Get(data interface{}) error {
	defaultProvider()
	return provider.Get(data)
}

func Set(data interface{}) (int, error) {
	defaultProvider()
	return provider.Set(data)
}
