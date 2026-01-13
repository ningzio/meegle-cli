# Meegle TUI Agent Guide

## Overview
This repository contains a Go-based TUI (Bubble Tea) application. The architecture follows MVU and separates routing, screens, store/reducers, and UI components.

## Development
- Use Go 1.22+.
- Keep screens isolated by using the `internal/screen` interfaces. Avoid importing `internal/app` from screens to prevent cycles.
- Business state updates should happen in `internal/store` reducers, not inside screen `Update` methods.
- New UI components should live under `internal/ui/components` and remain business-agnostic.

### Key directories
- `cmd/meegle-tui/`: application entrypoint.
- `internal/app/`: app model, router, overlays, keymaps, theme.
- `internal/screen/`: screen interfaces for decoupling.
- `internal/screens/`: screen implementations.
- `internal/store/`: domain state, messages, reducers.
- `internal/meegle/`: API client, auth, and command factory stubs.

## Coding Standards
- **Architecture**: Strictly follow Clean Architecture and MVU (Model-View-Update) principles.
- **Separation of Concerns**: Ensure logic is separated into correct layers (Entities -> Use Cases -> Adapters).
- **State Management**: All state mutations must occur within reducers (`internal/store`).

## Testing
- **Coverage**: All code changes must include tests. Aim for **100% test coverage**.
- **Execution**: Run tests using `go test ./...` before submitting.

## Linting
- **Tool**: Use `golangci-lint` version **v2.7+**.
- **Configuration**: The project uses a comprehensive `.golangci.yml`. ensure all linters pass.
- **Command**: Run `golangci-lint run` to verify your code.

## Design Guidelines
- State mutations only in reducers (`internal/store/reducers.go`).
- Screens should dispatch commands and read state via the `screen.AppModel` interface.
- Keep router transitions in `internal/app/router.go`.
- Use Lip Gloss for styling and keep theme tokens in `internal/app/theme.go`.
