# Meegle CLI Agent Instructions

This project is a Terminal User Interface (TUI) for managing Meegle tasks.
It is built using Golang and the Bubble Tea framework (https://github.com/charmbracelet/bubbletea).

## Project Goals
- Implement a TUI for Meegle task management.
- Follow the Model-View-Update (MVU) architecture provided by Bubble Tea.
- Ensure high code quality and test coverage.

## Architecture
- **Clean Architecture**:
  - `cmd/`: Entry points.
  - `internal/model/`: Domain entities.
  - `internal/service/`: Business logic.
  - `internal/adapter/`: Infrastructure (API clients, etc.).
  - `internal/tui/`: Presentation layer (Bubble Tea models and components).

## Development Guidelines
- **MVU Pattern**: All TUI components should strictly follow the Bubble Tea MVU pattern.
- **Testing**: All code must be covered by unit tests.
- **Linting**: Code must pass `golangci-lint` with the configuration in `.golangci.yml`.
- **API**: Currently using a Mock API. Real implementation will follow.

## Tech Stack
- Go (Golang)
- Bubble Tea (TUI Framework)
- Cobra (CLI Framework - optional, but good for entry)
- Viper (Configuration - optional)
