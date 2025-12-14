# nuke-port

A simple, cross-platform CLI tool to kill processes listening on a specific port.

## Features

- **Cross-platform**: Works on macOS, Linux, and Windows.
- **Interactive**: Asks for confirmation before killing a process.
- **Fast**: Written in Go for instant execution.
- **Force Mode**: Option to skip confirmation.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap geekaara/tap
brew install nuke-port
```

### Windows (Winget / Scoop / Manual)

**Manual**: Download the `.zip`/`.exe` from [Releases](https://github.com/geekaara/nuke-port/releases).

**Scoop** (Coming Soon):

```bash
scoop install https://github.com/geekaara/nuke-port/releases/latest/download/nuke-port_Windows_x86_64.zip
```

_(Once releases are generated)_

### From Source

```bash
git clone https://github.com/geekaara/nuke-port.git
cd nuke-port
go build -o nuke
```

## Usage

**Interactive Mode** (Safest):

```bash
./nuke 8080
```

_Will prompt: `Found process(es) [...] listening on port 8080. Kill them? (y/N):`_

**Force Mode** (Non-interactive):

```bash
./nuke --force 8080
# or shorthand
./nuke -f 8080
```

## Requirements

- **macOS/Linux**: Requires `lsof` command.
- **Windows**: Requires standard `netstat` and `taskkill` commands.
