# Meegle Task Management TUI

An MVP terminal UI for managing Meegle tasks using Bubble Tea + Bubbles + Lipgloss.

## Run

```bash
go run ./cmd/meegle-tui
```

Optional helpers:

```bash
make run
```

## Configuration

By default the app uses a **mock client** so it runs without setup. To use the real API, set:

```bash
export MEEGLE_BASE_URL="https://api.meegle.ai"
export MEEGLE_PLUGIN_ID="your-plugin-id"
export MEEGLE_PLUGIN_SECRET="your-plugin-secret"
export MEEGLE_PROJECT_KEY="your-project-key"
export MEEGLE_USER_KEY="your-user-key"
```

If any of the variables are missing, the app falls back to mock data.

> Note: The real client is wired for extension but returns "not implemented" until API details are filled in.

## Keymap

**Global**
- `q` / `ctrl+c`: quit

**Task list**
- `↑` / `↓`: move selection
- `enter`: open task detail
- `n`: new task
- `a`: add subtask to selected task
- `r`: refresh

**Task detail**
- `space`: complete/reopen selected subtask (with confirm)
- `a`: add subtask
- `esc`: back

**Editor**
- `enter`: save
- `esc`: back

## Structure

```
cmd/meegle-tui
internal/app
internal/store
internal/screens/{tasks,detail,editor}
internal/ui/components/{toast,modal}
internal/meegle
```
