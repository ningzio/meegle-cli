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

## Testing & Checks
Run these before every commit:
- `golangci-lint run`
- `go build ./...`

## Design Guidelines
- State mutations only in reducers (`internal/store/reducers.go`).
- Screens should dispatch commands and read state via the `screen.AppModel` interface.
- Keep router transitions in `internal/app/router.go`.
- Use Lip Gloss for styling and keep theme tokens in `internal/app/theme.go`.
