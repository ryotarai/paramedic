package outputlog

import (
	"fmt"
	"io"
	"time"
)

type Printer struct {
	Writer           io.Writer
	InstanceIDToName map[string]string
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		Writer: writer,
	}
}

func (p *Printer) Print(events []*Event) {
	for _, e := range events {
		fmt.Fprintf(p.Writer, "[%s] [%s] %s\n", e.Timestamp.Format(time.RFC3339), e.InstanceID(), e.Message)
	}
}
