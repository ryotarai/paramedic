package outputlog

import (
	"sync"

	"github.com/fatih/color"
)

var allColors = []*color.Color{
	color.New(color.FgGreen),
	color.New(color.FgYellow),
	color.New(color.FgBlue),
	color.New(color.FgMagenta),
	color.New(color.FgCyan),
	color.New(color.FgRed),
}

type Colorer struct {
	colors map[string]*color.Color
	mutex  sync.Mutex
}

func NewColorer() *Colorer {
	return &Colorer{
		colors: map[string]*color.Color{},
		mutex:  sync.Mutex{},
	}
}

func (c *Colorer) Color(key string) *color.Color {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	color, ok := c.colors[key]
	if !ok {
		color = allColors[len(c.colors)%len(allColors)]
		c.colors[key] = color
	}
	return color
}
