package log

import "testing"

func TestPrint(t *testing.T) {

	Init(MODE_LOCALFILE, "./temp.txt")

	Print("String=%s Int=%d", "hello world", 666)
	Info("String=%s Int=%d", "hello world", 666)
	Notice("String=%s Int=%d", "hello world", 666)
	Warning("String=%s Int=%d", "hello world", 666)
	Error("String=%s Int=%d", "hello world", 666)
	Critical("String=%s Int=%d", "hello world", 666)
	Alert("String=%s Int=%d", "hello world", 666)
	Emergency("String=%s Int=%d", "hello world", 666)
	// Fatal("String=%s Int=%d", "hello world", 666)
	Debug("String=%s Int=%d", "hello world", 666)
	Trace("String=%s Int=%d", "hello world", 666)
	TraceVerbose("String=%s Int=%d", "hello world", 666)

	if err := GetLastError(); err != nil {
		t.Errorf("last error in log package=%s", err.Error())
	}

}
