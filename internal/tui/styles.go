package tui

import "github.com/charmbracelet/lipgloss"

var titleStyle = lipgloss.NewStyle().
	Bold(true)

var selectedServiceStyle = lipgloss.NewStyle().
	Bold(true)

var servicePanelStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	Width(30).
	Height(14)

var detailPanelStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	Width(60).
	Height(14)

var findingHighStyle = lipgloss.NewStyle().
	Bold(true)

var findingMediumStyle = lipgloss.NewStyle()
