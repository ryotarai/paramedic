package outputlog

import (
	"sort"
	"strings"
	"time"
)

type Event struct {
	Message   string
	Timestamp time.Time
	LogStream string
}

func (e *Event) InstanceID() string {
	parts := strings.Split(e.LogStream, "/")
	return parts[len(parts)-1]
}

func SortEventsByTimestamp(events []*Event) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})
}
