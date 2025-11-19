# TaskGo Architecture

TaskGo is designed with modularity and scalability in mind.

## Directory Structure

```
taskgo/
├── cmd/            # Cobra commands (root, add, list, etc.)
├── internal/
│   ├── task/       # Task model and Manager logic
│   ├── storage/    # Storage interface and JSON implementation
│   └── ui/         # Lipgloss styles and UI helpers
├── docs/           # Documentation
├── main.go         # Entry point
└── go.mod          # Go module definition
```

## Components

### Task Manager (`internal/task`)
The `Manager` struct encapsulates the business logic for managing tasks. It relies on the `Storage` interface for data persistence, making it easy to swap out the storage backend (e.g., to SQLite or a remote API) without changing the core logic.

### Storage (`internal/storage`)
The `JSONStorage` implementation handles reading and writing tasks to a JSON file located at `~/.taskgo/tasks.json`. It ensures thread-safe access (though currently the CLI is single-threaded per invocation).

### UI (`internal/ui`)
We use [Lipgloss](https://github.com/charmbracelet/lipgloss) for styling. All styles are defined centrally in `style.go` to maintain consistency across the application.

### CLI (`cmd/`)
[Cobra](https://github.com/spf13/cobra) is used for command routing and flag parsing. Each command is defined in its own file for better maintainability.
