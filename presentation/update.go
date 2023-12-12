package presentation

import (
	"github.com/KarnerTh/plogger/extract"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case string:
		// TODO: remove
		m.logLines.Push(msg)
	case extract.Data:
		m.logLines.Push(msg.LogLine)
		m.values = append(m.values, msg.Value)
	}

	return m, nil
}
