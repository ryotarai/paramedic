package outputlog

import (
	"fmt"
	"io"
)

type Printer struct {
	Writer io.Writer

	colorer *Colorer
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		Writer:  writer,
		colorer: NewColorer(),
	}
}

func (p *Printer) Print(events []*Event) {
	resetColor := "\x1b[0m"

	for _, e := range events {
		instance := e.InstanceID()
		instance = p.colorer.Color(instance).Sprint(instance)

		fmt.Fprintf(p.Writer,
			"%s%s | %s | %s\n",
			resetColor,
			e.Timestamp.Format("15:04:05"),
			instance,
			e.Message)
	}
}
