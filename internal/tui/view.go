package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Compose Inspector\n\n")

	var servicesBuilder strings.Builder
	var detailsBuilder strings.Builder

	for i, service := range m.Services {
		cursor := "  "

		line := cursor + service

		if i == m.Cursor {
			line = selectedServiceStyle.Render("> " + service)
		}

		servicesBuilder.WriteString(line + "\n")

	}

	selectedName := m.Services[m.Cursor]
	selectedService := m.ComposeFile.Services[selectedName]

	detailsBuilder.WriteString("Selected Service: " + selectedName + "\n")

	// if containerName
	if selectedService.ContainerName != "" {
		detailsBuilder.WriteString("  Container Name: " + selectedService.ContainerName + "\n")
	}

	// if Image
	if selectedService.Image != "" {
		detailsBuilder.WriteString("  Image: " + selectedService.Image + "\n")
	}

	// if BuildContext
	if selectedService.Build.Context != "" {
		detailsBuilder.WriteString("  Build context : " + selectedService.Build.Context + "\n")
		detailsBuilder.WriteString("  Build dockerfile : " + selectedService.Build.Dockerfile + "\n")
	}

	// if ports
	if selectedService.Ports != nil && len(selectedService.Ports) > 0 {
		detailsBuilder.WriteString("  Ports : " + strings.Join(selectedService.Ports, ", ") + "\n")
	}

	// if command
	command := selectedService.Command
	if command != "" {
		if len(command) > 60 {
			command = command[:57] + "..."
		}

		detailsBuilder.WriteString("  Command: " + command + "\n")
	}

	// if networks
	if selectedService.Networks != nil && len(selectedService.Networks) > 0 {
		detailsBuilder.WriteString("  Networks : " + strings.Join(selectedService.Networks, ", ") + "\n")
	}

	detailsBuilder.WriteString("\nFindings:\n")

	found := false

	for _, finding := range m.Findings {
		if finding.Service != selectedName {
			continue
		}

		found = true

		line := finding.Severity + " - " + finding.Message

		if finding.Severity == "high" {
			line = findingHighStyle.Render(line)
		} else if finding.Severity == "medium" {
			line = findingMediumStyle.Render(line)
		}

		detailsBuilder.WriteString("  " + line + "\n")
	}

	if !found {
		detailsBuilder.WriteString("  No findings\n")
	}

	leftPanel := servicePanelStyle.Render(servicesBuilder.String())
	rightPanel := detailPanelStyle.Render(detailsBuilder.String())

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		" ",
		rightPanel,
	)

	b.WriteString(content)

	b.WriteString("\n↑/↓ navigate • q quit\n")

	return b.String()
}
