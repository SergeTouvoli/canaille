package main

import (
	"fmt"
	"os"

	"github.com/SergeTouvoli/canaille/internal/analysis"
	"github.com/SergeTouvoli/canaille/internal/compose"
	"github.com/SergeTouvoli/canaille/internal/tui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: canaille <compose-file>")
		os.Exit(1)
	}
	//fmt.Println("Compose file : " + os.Args[1])

	composeFile, err := compose.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing compose file: " + err.Error())
		os.Exit(1)
	}

	findings := analysis.Analyze(composeFile)

	//tui.PrintSummary(composeFile, findings)

	err = tui.Run(composeFile, findings)

	if err != nil {
		fmt.Println("Error starting TUI:", err)
		os.Exit(1)
	}
}
