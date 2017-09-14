package outputlog

import (
	"fmt"
	"io"
	"time"
)

type Printer struct {
	writer io.Writer
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		writer: writer,
	}
}

func (p *Printer) Print(events []*Event) {
	for _, e := range events {
		fmt.Fprintf(p.writer, "[%s] [%s] %s\n", e.Timestamp.Format(time.RFC3339), e.InstanceID(), e.Message)
	}
}
