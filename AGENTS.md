# Meegle CLI Developer Guide

## Project overview
- Bubble Tea-based TUI for managing Meegle tasks.
- MVU architecture with routing and store separation.

## Development conventions
- Keep screens isolated as independent packages and have each implement the `app.Screen` interface.
- All I/O must be done via `tea.Cmd`; never mutate state inside Cmds.
- Use request correlation (`reqID`) for asynchronous operations.
- Prefer small, focused files and avoid giant `Update` switches.

## Quality gates (required before every commit)
- Run `golangci-lint` and fix lint errors:
  ```bash
  golangci-lint run ./...
  ```
- Ensure unit tests maintain **>= 80% coverage**:
  ```bash
  go test ./... -cover
  ```

If these commands cannot be run due to environment limitations, record the failure in the PR/testing notes.
