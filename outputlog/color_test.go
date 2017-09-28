package outputlog

import (
	"testing"

	"github.com/fatih/color"
)

func TestColorer(t *testing.T) {
	examples := []struct {
		key   string
		color *color.Color
	}{
		{key: "a", color: color.New(color.FgGreen)},
		{key: "b", color: color.New(color.FgYellow)},
		{key: "a", color: color.New(color.FgGreen)},
		{key: "c", color: color.New(color.FgBlue)},
	}

	c := NewColorer()
	for _, e := range examples {
		if !c.Color(e.key).Equals(e.color) {
			t.Errorf("color for %s is wrong", e.key)
		}
	}
}
