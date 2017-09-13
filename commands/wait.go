package commands

import (
	"log"
	"time"

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type WaitStatusOptions struct {
	SSM   awsclient.SSM
	Store *store.Store

	CommandID string
	Statuses  []string
	Interval  time.Duration
}

func WaitStatus(opts *WaitStatusOptions) chan *Command {
	interval := opts.Interval
	if interval == time.Duration(0) {
		interval = 15 * time.Second
	}

	c := make(chan *Command)

	go func() {
		for {
			log.Printf("[DEBUG] Checking status of command %s", opts.CommandID)
			command, err := Get(&GetOptions{
				SSM:       opts.SSM,
				Store:     opts.Store,
				CommandID: opts.CommandID,
			})
			if err != nil {
				log.Printf("[WARN] %s", err)
			}

			for _, st := range opts.Statuses {
				if command.Status == st {
					c <- command
					return
				}
			}

			time.Sleep(interval)
		}
	}()

	return c
}
