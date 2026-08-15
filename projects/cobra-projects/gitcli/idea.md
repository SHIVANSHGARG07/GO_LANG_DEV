gitcli init                           # Initialize a repo
gitcli status                         # Show status
gitcli commit -m "message"            # Create a commit
gitcli log                            # Show commit history
gitcli remote add <name> <url>        # Add remote (nested!)
gitcli remote list                    # List remotes (nested!)
gitcli remote remove <name>           # Remove remote (nested!)
gitcli config --global user.name "X"  # Set config
gitcli config --list                  # List config




gitcli/
├── main.go
├── go.mod
├── cmd/
│   ├── root.go          # Root command + persistent flags
│   ├── init.go          # Initialize repo
│   ├── status.go        # Show status
│   ├── commit.go        # Create commit
│   ├── log.go           # Show history
│   ├── remote.go        # Remote parent command
│   ├── remote_add.go    # Nested: remote add
│   ├── remote_list.go   # Nested: remote list
│   ├── remote_remove.go # Nested: remote remove
│   └── config.go        # Config management
├── internal/
│   ├── repo.go          # Repository management
│   ├── commit.go        # Commit storage
│   └── config.go        # Config file handling
└── .gitcli/             # Created by 'gitcli init'
    ├── commits.json
    ├── config.json
    └── remotes.json