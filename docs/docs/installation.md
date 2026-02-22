# Installation

Get started with **go-microx** by installing the CLI on your machine.

## Prerequisites

Before installing, ensure you have the following requirements met:

- **Go**: Version 1.20 or later installed. [Download Go](https://go.dev/dl/)
- **Git**: Installed and configured.
- **Docker**: (Optional) For running the generated services in containers.

## Install via Go

The easiest way to install **go-microx** is using the `go install` command:

```bash
go install github.com/abhishekkushwahaa/go-microx/cmd/go-microx@latest
```

Ensure your `$GOPATH/bin` is in your system's `PATH` to run the command globally.

## Verify Installation

Check if the CLI is installed correctly by running:

```bash
go-microx --version
```

If everything is set up correctly, you should see the current version of the tool.

## Updating

To update to the latest version, simply run the installation command again:

```bash
go install github.com/abhishekkushwahaa/go-microx/cmd/go-microx@latest
```
