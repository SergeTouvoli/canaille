package tui

import (
	"fmt"
	"strings"

	"github.com/SergeTouvoli/canaille/internal/analysis"
	"github.com/SergeTouvoli/canaille/internal/compose"
)

func PrintSummary(composeFile *compose.ComposeFile, findings []analysis.Finding) {

	for _, finding := range findings {
		fmt.Printf("Service: %s, Message: %s, Severity: %s\n", finding.Service, finding.Message, finding.Severity)
	}

	for key, service := range composeFile.Services {
		fmt.Println("Service : " + key)

		// Si précence de image:
		if service.Image != "" {
			fmt.Println("Image : " + service.Image)
		}

		// si précende de ports
		if service.Ports != nil && len(service.Ports) > 0 {
			for _, port := range service.Ports {
				fmt.Println("Port : " + port)
			}
		}

		// si précense de commands
		if service.Command != "" {
			fmt.Println("Command : " + service.Command)
		}

		// Si précende de build:
		if service.Build.Context != "" {
			fmt.Println("Build context : " + service.Build.Context)
			fmt.Println("Build dockerfile : " + service.Build.Dockerfile)
		}

		// si présence de container_name:
		if service.ContainerName != "" {
			fmt.Println("Container name : " + service.ContainerName)
		}

		// si précende de restart:
		if service.Restart != "" {
			fmt.Println("Restart : " + service.Restart)
		}

		// si précense de networks
		if service.Networks != nil && len(service.Networks) > 0 {
			fmt.Println("Networks : " + strings.Join(service.Networks, ", "))
		}

		if len(service.Healthcheck.Test) > 0 {
			fmt.Println("Healthcheck : yes ")
		}

		if service.Healthcheck.Interval != "" {
			fmt.Println("Healthcheck interval : " + service.Healthcheck.Interval)
		}

		fmt.Println("-----")
	}
}
