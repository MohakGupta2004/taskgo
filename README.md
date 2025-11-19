# TaskGo

TaskGo is a production-grade, open-source CLI Todo List application written in Go. It features a beautiful table-based UI, a built-in Pomodoro timer, and robust task management capabilities. Designed for developers who live in the terminal.

## Features

- **Task Management**: Add, list, update, and remove tasks with ease.
- **Beautiful UI**: Colorful table output and banners using Lipgloss.
- **Pomodoro Timer**: Integrated focus timer to boost productivity.
- **Persistent Storage**: Tasks are saved locally in `~/.taskgo/tasks.json`.
- **Cross-Platform**: Works on Linux, macOS, and Windows.

## Installation

### Quick Install (Linux/macOS)

You can install TaskGo using the provided installation script:

```bash
curl -sL https://raw.githubusercontent.com/MohakGupta2004/taskgo/main/install.sh | sudo bash
```

### Manual Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/MohakGupta2004/taskgo.git
   cd taskgo
   ```

2. Build the binary using Make:
   ```bash
   make build
   ```
   This will create the binary in `bin/taskgo`.

3. Install globally:
   ```bash
   make install
   ```

## Usage

### Add a Task
```bash
taskgo add "Buy groceries"
taskgo add "Finish Go project"
```

### List Tasks
```bash
taskgo list
```

### Update Task Status
Status options: `todo`, `in-progress`, `completed`
```bash
taskgo update 1 in-progress
taskgo update 1 completed
```

### Remove a Task
```bash
taskgo remove 1
```

### Start Pomodoro Timer
Default duration is 25 minutes.
```bash
taskgo pomodoro
taskgo pomodoro 45  # Start for 45 minutes
```

## Architecture

TaskGo follows a clean architecture pattern:
- **cmd/**: CLI command definitions using Cobra.
- **internal/task/**: Core business logic and models.
- **internal/storage/**: Data persistence layer.
- **internal/ui/**: UI styling and rendering.

For more details, check the [docs/](docs/) folder.

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md).

## License

MIT License
