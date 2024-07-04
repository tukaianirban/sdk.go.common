package log

import (
	"testing"
)

func TestPrint(t *testing.T) {

	Init(true)

	type DummyPerson struct {
		Name   string
		Income float64
	}
	tom := DummyPerson{
		Name:   "tom",
		Income: 9.84,
	}

	Info("info about Tom = %#v", tom)
	Warning("this should be a warning mesage with string=%s Int=%d", "hello world", 121)
	Error("error flagged with string=%s Int=%d", "hello world", 363)
	Alert("this is an alert for string=%s Int=%d", "hello world", 424)
	Debug("this log contains debugging info of string=%s Int=%d", "hello world", 7770)

	if err := GetLastError(); err != nil {
		t.Errorf("last error in log package=%s", err.Error())
	}
}
