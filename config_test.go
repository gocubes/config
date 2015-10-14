package config

import (
	"os"
	"testing"
	"time"
)

var (
	jsonFileName string = time.Now().Format("200601020304.json")
	jsonProvider Provider
)

type TestData struct {
	Date time.Time
	Name string
}

func TestJsonSet(t *testing.T) {
	testData := &TestData{
		Date: time.Now(),
		Name: "default",
	}

	// new provider
	jsonProvider, err := New(jsonFileName, "json")
	if err != nil {
		// t.Logf("[JSON New]%v", err.Error())
		t.Fatalf("[JSON New]%v", err.Error())
		return
	}

	n, err := jsonProvider.Set(*testData)
	if err != nil {
		t.Fatalf("[JSON]write %v bytes, error: %v\n", n, err)
	}
}

func TestCleanCase(t *testing.T) {
	t.Logf("Remove json file: %s\n", jsonFileName)
	os.Remove(jsonFileName)
}
