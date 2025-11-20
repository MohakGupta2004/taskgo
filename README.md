# TaskGo

TaskGo is a production-grade, open-source CLI Todo List application written in Go. It features a beautiful table-based UI, a built-in Pomodoro timer, and robust task management capabilities. Designed for developers who live in the terminal.

## Features

- **Task Management**: Add, list, update, edit, and remove tasks with ease.
- **Grouped Tasks**: Organize tasks into groups (e.g., "Work", "Personal") with a tree-view.
- **Context Switching**: "Checkout" a group to automatically add tasks to it.
- **Task Validity & Auto-Removal**: Set expiration times for tasks - they auto-remove when expired.
- **Group Validity Defaults**: Configure default validity periods per group.
- **Unquoted Input**: Add tasks and set validity without quotation marks.
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

**Basic usage** (no quotes needed):
```bash
taskgo add Buy groceries
```

**Add to a specific group:**
```bash
taskgo add Finish Go project -g work
```

**Add with validity duration** (positional argument):
```bash
taskgo add 2h Complete report
# Task will auto-remove after 2 hours
```

**Add with validity flag:**
```bash
taskgo add "Team meeting" -v 30m -g work
```

Supported duration formats: `10s`, `5m`, `2h`, `24h`

### Edit a Task

Update the title of an existing task (no quotes needed):
```bash
taskgo edit 1 New task title here
```

### List Tasks

Tasks are displayed in a tree structure, grouped by their category. The list shows:
- Task ID, Title, Status
- Created At, Completed At timestamps
- Valid Until (remaining time or "Expired")

```bash
taskgo list
```

Expired tasks are automatically removed when you run `list`.

### Task Groups & Context

**Checkout a group:**
```bash
taskgo checkout work
taskgo add Meeting notes  # Added to 'work' group automatically
```

**Configure group validity** (positional argument):
```bash
taskgo group 8h work
# All tasks added to 'work' will default to 8h validity
```

**Configure with flag:**
```bash
taskgo group personal -v 24h
```

**List all groups:**
```bash
taskgo group list
```

**Show current group:**
```bash
taskgo group
```

### Update Task Status

Status options: `todo`, `in-progress`, `completed`
```bash
taskgo update 1 in-progress
taskgo update 1 completed
```

### Remove a Task

Remove a single task by ID:
```bash
taskgo remove 1
```

Remove **ALL** tasks in the current group:
```bash
taskgo remove all
# OR
taskgo remove "*"
```

### Upgrade TaskGo

Update the executable to the latest version from the repository:
```bash
taskgo upgrade
```

### Start Pomodoro Timer

Default duration is 25 minutes.
```bash
taskgo pomodoro              # 25 minutes
taskgo pomodoro 45           # 45 minutes
taskgo pomodoro 01:30:00     # 1 hour 30 minutes
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
