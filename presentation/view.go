package presentation

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func (m model) View() string {
	var out string
	for _, item := range m.logLines {
		out += fmt.Sprintf("%s\n", item)
	}

	if len(m.values) > 0 {
		graph := asciigraph.Plot(m.values, asciigraph.Height(10))
		out += "\n\n"
		out += graph
		out += "\n\n"
	}

	return out
}
