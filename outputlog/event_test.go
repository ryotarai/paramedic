package outputlog

import (
	"testing"
	"time"
)

func TestSortEventsByTimestamp(t *testing.T) {
	events := []*Event{
		{Message: "foo", Timestamp: time.Unix(1, 0)},
		{Message: "bar", Timestamp: time.Unix(0, 0)},
		{Message: "baz", Timestamp: time.Unix(2, 0)},
	}
	expected := []*Event{
		events[1],
		events[0],
		events[2],
	}

	SortEventsByTimestamp(events)

	for i, e := range expected {
		if events[i] != e {
			t.Errorf("sorting events failed, expected %+v but got %+v", e, events[i])
		}
	}
}

func TestEventInstanceID(t *testing.T) {
	e := &Event{
		LogStream: "foo/bar/baz",
	}
	if got := e.InstanceID(); got != "baz" {
		t.Errorf("Event.InstanceID() = %v, want %v", got, "baz")
	}
}
