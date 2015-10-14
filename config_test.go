package config

import (
	"os"
	"testing"
	"time"
)

var (
	name         = "default"
	now          = time.Now()
	jsonFileName = now.Format("200601020304.json")
)

type TestData struct {
	Date time.Time
	Name string
}

func TestJsonSet(t *testing.T) {
	testData := &TestData{
		Date: now,
		Name: name,
	}

	// new provider
	jsonProvider, err := New(jsonFileName, "json")
	if err != nil {
		t.Fatalf("[JSON New]%v", err.Error())
	}

	// set data
	n, err := jsonProvider.Set(*testData)
	if err != nil {
		t.Fatalf("[JSON Set]write %v bytes, error: %v\n", n, err)
	}
}

func TestJsonGet(t *testing.T) {
	testData := &TestData{}

	// new provider
	jsonProvider, err := New(jsonFileName, "json")
	if err != nil {
		t.Fatalf("[JSON New]%v", err.Error())
	}

	// get data
	err = jsonProvider.Get(testData)
	if err != nil {
		t.Fatalf("[JSON Get] %v\n", err.Error())
	}

	// check data
	if testData.Date != now || testData.Name != name {
		t.Fatalf("[Check]set Date: %v, Name: %v;but get Date: %v, Name: %v\n", now, name, testData.Date, testData.Name)
	}
}

func TestCleanCase(t *testing.T) {
	t.Logf("Remove json file: %s\n", jsonFileName)
	os.Remove(jsonFileName)
}
