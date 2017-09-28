package outputlog

import (
	"bytes"
	"testing"
	"time"
)

func TestPrinter(t *testing.T) {
	writer := bytes.NewBufferString("")

	events := []*Event{
		{Message: "foo", Timestamp: time.Unix(0, 0).UTC(), LogStream: "foo/i-aaa"},
		{Message: "bar", Timestamp: time.Unix(1, 0).UTC(), LogStream: "foo/i-bbb"},
	}

	p := NewPrinter(writer)
	p.Print(events)

	expects := []string{
		"00:00:00 | i-aaa | foo\n",
		"00:00:01 | i-bbb | bar\n",
	}

	for _, e := range expects {
		l, err := writer.ReadString(byte('\n'))
		if err != nil {
			t.Error(err)
		}

		if l != e {
			t.Errorf("got %v, want %v", l, e)
		}
	}
}
