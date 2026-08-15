# 🚀 GitCLI - A Git-Inspired CLI Tool

A simplified Git-like command-line interface built with Go and Cobra, demonstrating CLI development concepts including nested commands, persistent flags, and JSON-based data storage.

## 📚 Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Commands](#commands)
- [Project Structure](#project-structure)
- [Examples](#examples)
- [Learning Goals](#learning-goals)
- [Technical Details](#technical-details)
- [Contributing](#contributing)
- [License](#license)

---

## ✨ Features

- 📦 **Repository Management** - Initialize and manage local repositories
- 💾 **Commit System** - Create and track commits with messages and authors
- 📜 **Commit History** - View commit logs in reverse chronological order
- 🌐 **Remote Management** - Add, list, and remove remote repositories
- 🎯 **Persistent Flags** - Global flags available across all commands
- 💻 **Clean Architecture** - Separation of CLI logic (`cmd/`) and business logic (`internal/`)

---

## 📥 Installation

### Prerequisites

- Go 1.21 or higher
- Git (for cloning the repository)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/SHIVANSHGARG07/gitcli.git
cd gitcli

# Install dependencies
go mod download

# Build the binary
go build -o gitcli

# (Optional) Add to PATH
sudo mv gitcli /usr/local/bin/




########### Quick Start ######

# Initialize a new repository
./gitcli init

# Create your first commit
./gitcli commit -m "Initial commit" -a "Your Name"

# View commit history
./gitcli log

# Add a remote
./gitcli remote add origin https://github.com/user/repo

# List remotes
./gitcli remote list





######## Init Repo ########
./gitcli init

# With verbose output
./gitcli init -v




######### Commit ########
.# Basic commit
./gitcli commit -m "Add new feature"

# Commit with author
./gitcli commit -m "Fix bug" -a "John Doe"

# Using verbose mode
./gitcli commit -m "Update docs" -v




######### Log ########
# Basic log
./gitcli log

# Using verbose mode
./gitcli log -v



######### Remote ########
# Add a remote
./gitcli remote add origin https://github.com/user/repo

# List remotes
./gitcli remote list

# Remove a remote
./gitcli remote remove origin





gitcli/
├── cmd/                    # CLI command definitions
│   ├── root.go            # Root command & global flags
│   ├── init.go            # Init command
│   ├── commit.go          # Commit command
│   ├── log.go             # Log command
│   └── remote_cmd/        # Remote command group
│       ├── remote.go      # Parent command
│       ├── remote_add.go
│       ├── remote_list.go
│       └── remote_remove.go
├── internal/              # Business logic
│   ├── init.go           # Repository initialization
│   ├── commit.go         # Commit management
│   ├── remote.go         # Remote management
│   └── util.go           # Utility functions
├── main.go               # Application entry point
├── go.mod                # Go module definition
└── README.md             # This file