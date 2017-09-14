package outputlog

import (
	"time"
)

type Reader interface {
	Read() ([]*Event, error)
}

func Follow(r Reader, p *Printer, stopCh chan struct{}) error {
	exit := false
	for {
		events, err := r.Read()
		if err != nil {
			return err
		}
		p.Print(events)

		if exit {
			break
		}

		select {
		case <-stopCh:
			exit = true
		case <-time.After(10 * time.Second):
		}
	}

	return nil
}
