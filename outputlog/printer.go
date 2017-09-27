package outputlog

import (
	"fmt"
	"io"
)

type Printer struct {
	Writer           io.Writer
	InstanceIDToName map[string]string

	colorer *Colorer
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		Writer:  writer,
		colorer: NewColorer(),
	}
}

func (p *Printer) Print(events []*Event) {
	for _, e := range events {
		fmt.Fprintf(p.Writer,
			"%s | %s | %s\n",
			e.Timestamp.Format("15:04:05"),
			p.colorer.Color(e.InstanceID()).Sprint(e.InstanceID()),
			e.Message)
	}
}
