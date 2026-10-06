package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SergeTouvoli/canaille/internal/analysis"
	"github.com/SergeTouvoli/canaille/internal/compose"
	"github.com/SergeTouvoli/canaille/internal/tui"
)

func main() {

	checkMode := flag.Bool("check", false, "run in non-interactive check mode")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: canaille <compose-file>")
		os.Exit(1)
	}

	composeFile, err := compose.ParseFile(flag.Args()[0])
	if err != nil {
		fmt.Println("Error parsing compose file: " + err.Error())
		os.Exit(1)
	}

	findings := analysis.Analyze(composeFile)

	if *checkMode {

		if len(findings) > 0 {
			fmt.Println("Findings:")
			for _, finding := range findings {
				fmt.Printf("- Service: %s, Title: %s, Severity: %s\n", finding.Service, finding.Title, finding.Severity)
				if finding.Description != "" {
					fmt.Printf("  Description: %s\n", finding.Description)
				}
			}
			os.Exit(1)
		} else {
			fmt.Println("No findings.")
			os.Exit(0)
		}
	}

	err = tui.Run(composeFile, findings)

	if err != nil {
		fmt.Println("Error starting TUI:", err)
		os.Exit(1)
	}
}
