package tui

import (
	"sort"

	"github.com/SergeTouvoli/canaille/internal/analysis"
	"github.com/SergeTouvoli/canaille/internal/compose"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(composeFile *compose.ComposeFile, findings []analysis.Finding) error {
	var services []string

	for name := range composeFile.Services {
		services = append(services, name)
	}

	sort.Strings(services)

	model := Model{
		Services:    services,
		ComposeFile: composeFile,
		Findings:    findings,
		Cursor:      0,
	}

	program := tea.NewProgram(model)

	_, err := program.Run()

	return err
}
