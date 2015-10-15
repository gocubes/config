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
	xmlFileName  = now.Format("200601020304.xml")
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

func TestXmlSet(t *testing.T) {
	testData := &TestData{
		Date: now,
		Name: name,
	}

	// new provider
	xmlProvider, err := New(xmlFileName, "xml")
	if err != nil {
		t.Fatalf("[xml New]%v", err.Error())
	}

	// set data
	n, err := xmlProvider.Set(*testData)
	if err != nil {
		t.Fatalf("[xml Set]write %v bytes, error: %v\n", n, err)
	}
}

func TestXmlGet(t *testing.T) {
	testData := &TestData{}

	// new provider
	xmlProvider, err := New(xmlFileName, "xml")
	if err != nil {
		t.Fatalf("[xml New]%v", err.Error())
	}

	// get data
	err = xmlProvider.Get(testData)
	if err != nil {
		t.Fatalf("[xml Get] %v\n", err.Error())
	}

	// check data
	if !testData.Date.Equal(now) || testData.Name != name {
		t.Fatalf("[Check]set Date: %v, Name: %v;but get Date: %v, Name: %v\n", now, name, testData.Date, testData.Name)
	}
}

func TestCleanCase(t *testing.T) {
	t.Logf("Remove json file: %s\n", jsonFileName)
	os.Remove(jsonFileName)

	t.Logf("Remove xml file: %s\n", xmlFileName)
	os.Remove(xmlFileName)
}
