package tui

import (
	"github.com/SergeTouvoli/canaille/internal/analysis"
	"github.com/SergeTouvoli/canaille/internal/compose"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Services    []string
	ComposeFile *compose.ComposeFile
	Findings    []analysis.Finding
	Cursor      int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down":
			if m.Cursor < len(m.Services)-1 {
				m.Cursor++
			}
		}
	}

	return m, nil
}
