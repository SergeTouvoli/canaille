# 🐾 Canaille

> 🚧 **Work in Progress** — This project is currently under development.

**Canaille** is an interactive Docker Compose inspector and linter for the terminal, written in Go.

It parses Docker Compose files, lets you explore services through a terminal user interface, and detects potentially unsafe configurations and common bad practices.

The project is named after my Shih Tzu, Canaille. 🐶

## Why this project?

I'm currently learning Go and wanted to build a real project instead of only following tutorials or solving isolated exercises.

I already work with Docker and Docker Compose, so I decided to build a tool around something I use regularly while learning Go concepts such as:

- Structs and custom types
- Maps and slices
- Pointers
- Error handling
- Package organization
- YAML parsing
- Terminal user interfaces
- Separation between parsing, analysis, and presentation

This is both a learning project and an experiment toward building a useful Docker Compose inspection tool.

The code and architecture will evolve as I learn more about Go.

## Current Features

Canaille can currently:

- Parse a Docker Compose YAML file
- Detect and list services
- Display service information such as:
  - Image
  - Build context and Dockerfile
  - Container name
  - Published ports
  - Networks
  - Commands
  - Healthchecks
- Navigate between services using an interactive terminal interface
- Analyze Docker Compose configurations
- Display findings for the selected service

### Current Analysis Rules

Canaille currently detects:

- Images using the `latest` tag
- Services without an image or build context
- Ports published without an explicit host IP
- Ports explicitly bound to `0.0.0.0`

More rules will be added as the project evolves.

## Interface

Canaille uses:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the interactive terminal interface
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for terminal layout and styling
- [yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) for parsing Docker Compose files

Example:

```text
Canaille

╭──────────────────────────────╮ ╭────────────────────────────────────────────╮
│                              │ │                                            │
│    app                       │ │  Selected Service: nginx                   │
│    db                        │ │    Container Name: my_nginx                │
│    worker                    │ │    Image: nginx:latest                     │
│  > nginx                     │ │    Networks: default, proxy                │
│                              │ │                                            │
│                              │ │  Findings:                                 │
│                              │ │    medium - Service is using the           │
│                              │ │    'latest' tag for the image              │
╰──────────────────────────────╯ ╰────────────────────────────────────────────╯

↑/↓ navigate • q quit
```

## Usage

Run Canaille and provide the path to a Docker Compose file:

```bash
go run ./cmd/canaille ./compose.yaml
```

Navigation:

```text
↑ / ↓    Navigate between services
q        Quit
Ctrl+C   Quit
```

> The installation and distribution process will evolve as the project matures.

## Project Structure

```text
canaille/
├── cmd/
│   └── canaille/
│       └── main.go
├── internal/
│   ├── analysis/
│   ├── compose/
│   └── tui/
├── testdata/
├── go.mod
├── LICENSE
└── README.md
```

The project is separated into different responsibilities:

- `compose` — Docker Compose models and YAML parsing
- `analysis` — Configuration analysis and findings
- `tui` — Interactive terminal interface
- `cmd/canaille` — Application entry point

This separation should eventually allow the same analysis engine to be used outside the interactive TUI.

## Roadmap

Some ideas for future versions:

- More Docker Compose security checks
- Detect privileged containers
- Detect Docker socket mounts
- Analyze sensitive volume mounts
- Analyze healthchecks
- Analyze restart policies
- Detect potentially exposed services
- Improve network analysis
- Better finding severity levels
- Global Compose configuration score
- Improved TUI and navigation
- CI/CD friendly check mode
- JSON output
- Automated tests

A future non-interactive mode could make Canaille usable in CI/CD pipelines:

```bash
canaille compose.yaml --check
```

## Learning Project

I'm still a beginner in Go.

This repository is intentionally public so I can document my progress, practice writing idiomatic Go, experiment with project architecture, and improve the code over time.

Some implementations may therefore not yet be the most idiomatic or optimal Go solutions.

Feedback, suggestions, and code reviews are welcome.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.