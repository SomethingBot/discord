package logging

import (
	"bytes"
	"log"
	"testing"
)

func TestStandard_Log(t *testing.T) {
	buf := bytes.Buffer{}
	l := Standard{Logger: *log.New(&buf, "", 0)}
	l.Log(Info, "test")
	if buf.String() != "INFO: test\n" {
		t.Fatalf("logger did not log expected output (%v), got (%v)", "INFO: test", buf.String())
	}
}
