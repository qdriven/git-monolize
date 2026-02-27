# Monolize

A CLI application to manage multiple git repositories. Built with Go, Cobra, and Viper.

## Features

1. **Update all repositories** - Scan a directory and update all git repositories to their latest version
2. **Create mono repo** - Create a mono repository with all found repositories as submodules
3. **Sync submodules** - Update all submodules in a mono repo with a single git command

## Installation

```bash
go build -o monolize
```

## Usage

### Update all repositories

Update all git repositories in the current directory:

```bash
monolize update
```

Update repositories in a specific directory:

```bash
monolize update --path /path/to/repos
```

### Create a mono repo

Create a mono repo from all repositories in the current directory:

```bash
monolize create
```

Create with a custom name:

```bash
monolize create --name my-mono-repo
```

Create in a specific output path:

```bash
monolize create --output /path/to/output
```

### Sync submodules in a mono repo

Update all submodules to their latest versions:

```bash
monolize sync /path/to/mono-repo
```

## Configuration

You can create a configuration file at `~/.monolize.yaml`:

```yaml
path: /default/path/to/repos
default_branch: main
auto_commit: true
```

## Commands

| Command | Description |
|---------|-------------|
| `update` | Update all git repositories to the latest version |
| `create` | Create a mono repo with all repositories as submodules |
| `sync`   | Sync all submodules in a mono repo to the latest version |

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--path` | `-p` | Path to the directory containing git repositories |
| `--config` | | Config file path |
| `--name` | `-n` | Name of the mono repo directory (for create command) |
| `--output` | `-o` | Output path for the mono repo (for create command) |
