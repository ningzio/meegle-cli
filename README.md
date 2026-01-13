# meegle-cli
Manage your meegle project within CLI.

## Project Structure

- `cmd/`: Application entry points.
- `internal/app/`: Application model, routing, and global state.
- `internal/meegle/`: Meegle API client and integration.
- `internal/screen/`: Interface definitions for screens.
- `internal/screens/`: Implementations of various UI screens.
- `internal/store/`: State management (Redux/Elm-like pattern).
- `internal/ui/`: Reusable UI components.

## Requirements

- **Go**: 1.22+
- **golangci-lint**: v2.7+ (Use the latest v2 version for development)

## Development

Run linters:
```bash
golangci-lint run
```
