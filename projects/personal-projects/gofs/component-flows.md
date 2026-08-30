

1) How running main.go mounts to Mac

Your Go Code
    ↓
fs.Mount() function
    ↓
┌──────────────────────────────────────────────┐
│ Step A: Open FUSE Device                     │
├──────────────────────────────────────────────┤
│ - Connects to macFUSE kernel extension       │
│ - Opens /dev/osxfuse0 (or /dev/fuse)        │
│ - Gets a file descriptor                     │
└──────────────┬───────────────────────────────┘
               │
               ▼
┌──────────────────────────────────────────────┐
│ Step B: Send Mount Request to macOS         │
├──────────────────────────────────────────────┤
│ - Tells macOS: "I want to mount at           │
│   /Volumes/gofs"                             │
│ - Passes options (name, debug, etc.)         │
│ - Requests filesystem type: FUSE             │
└──────────────┬───────────────────────────────┘
               │
               ▼
┌──────────────────────────────────────────────┐
│ Step C: macOS Kernel Processes Request      │
├──────────────────────────────────────────────┤
│ 1. Checks if /Volumes/ exists (yes!)        │
│ 2. Creates /Volumes/gofs directory          │
│ 3. Marks it as a FUSE mount point           │
│ 4. Connects it to macFUSE kernel extension  │
└──────────────┬───────────────────────────────┘
               │
               ▼
┌──────────────────────────────────────────────┐
│ Step D: macFUSE Sets Up Communication       │
├──────────────────────────────────────────────┤
│ - Creates a channel between:                 │
│   macOS kernel ↔ macFUSE ↔ Your Go program  │
│ - Registers your RootNode as the handler    │
│ - Starts listening for requests              │
└──────────────┬───────────────────────────────┘
               │
               ▼
┌──────────────────────────────────────────────┐
│ Step E: Returns Server Object                │
├──────────────────────────────────────────────┤
│ - Mount successful!                          │
│ - Returns server object to your code         │
│ - Also returns err (nil if success)          │
└──────────────────────────────────────────────┘