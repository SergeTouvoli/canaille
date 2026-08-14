package analysis

import (
	"strings"

	"github.com/SergeTouvoli/canaille/internal/compose"
)

func Analyze(composeFile *compose.ComposeFile) []Finding {
	var findings []Finding

	for key, service := range composeFile.Services {

		if service.Image == "" && service.Build.Context == "" {
			findings = append(findings, Finding{
				Service:  key,
				Message:  "Service has no image or build context defined",
				Severity: "high",
			})
		} else if strings.HasSuffix(service.Image, ":latest") {
			findings = append(findings, Finding{
				Service:  key,
				Message:  "Service is using the 'latest' tag for the image",
				Severity: "medium",
			})
		}

		// Check for ports
		for _, port := range service.Ports {
			parts := strings.Split(port, ":")
			if len(parts) == 2 {
				findings = append(findings, Finding{
					Service:  key,
					Message:  "Published port has no explicit host IP and may be exposed on all interfaces",
					Severity: "medium",
				})
			} else if len(parts) == 3 {
				hostIP := parts[0]
				if hostIP == "0.0.0.0" {
					findings = append(findings, Finding{
						Service:  key,
						Message:  "Published port is exposed on all interfaces (0.0.0.0)",
						Severity: "medium",
					})
				}

			}
		}
	}

	return findings
}
