package bruce

import (
	"testing"
)

func TestGetString(t *testing.T) {

	Init(MODE_LOCAL, "./sample_config.json")

	val, err := GetString("firestore.collection")
	if err != nil {
		t.Fatalf("failed to get the firestore collection data, reason=%s", err.Error())
	}

	t.Logf("retrieved firestore collection data = %s", val)
}

func TestGetInt(t *testing.T) {

	if err := readFromFile("./sample_config.json"); err != nil {
		t.Fatalf("failed to init from config file, reason=%s", err.Error())
	}

	val, err := GetInt("firestore.fieldCount")
	if err != nil {
		t.Fatalf("failed to get the firestore fieldCount data, reason=%s", err.Error())
	}

	t.Logf("retrieved firestore fieldCount data = %d", val)

	val, err = GetInt("version")
	if err != nil {
		t.Fatalf("failed to get the version int")
	}

	t.Logf("retrieved version data = %d", val)
}

func TestGetFloat64(t *testing.T) {

	if err := readFromFile("./sample_config.json"); err != nil {
		t.Fatalf("failed to init from config file, reason=%s", err.Error())
	}

	val, err := GetFloat64("semanticVersion")
	if err != nil {
		t.Fatalf("failed to get the semanticVersion data, reason=%s", err.Error())
	}

	t.Logf("retrieved semanticVersion data = %f", val)
}
