# gobox sandbox

A minimal and cross-platform sandbox for executing external commands in Go.
It provides controlled execution with timeouts, working directory isolation,
environment variables, and structured results.

This library is designed to be simple, explicit, and close to the Go standard
library philosophy.

---

## Features

- Command execution with timeout
- Custom working directory
- Custom environment variables
- Separate stdout and stderr capture
- Exit code reporting
- Context-aware execution
- Cross-platform (Windows, Linux, macOS)

---

## Installation

```bash
go get github.com/yusnelgg/gobox
```

---

## Basic Usage

```go
package main

import (
	"context"
	"time"

	"github.com/yusnelgg/gobox/pkg/sandbox"
)

func main() {
	result, err := sandbox.Run(context.Background(), sandbox.Config{
		Path:    "go",
		Args:    []string{"version"},
		Timeout: 2 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	println("stdout:", result.Stdout)
	println("stderr:", result.Stderr)
	println("exit code:", result.ExitCode)
}
```

---

## Running Shell Commands

### Windows (`cmd.exe`)

Some commands like `echo`, `dir`, or `cd` are shell built-ins on Windows
and cannot be executed directly.

```go
result, err := sandbox.Run(context.Background(), sandbox.Config{
	Path:    "cmd",
	Args:    []string{"/C", "echo Hello, World!"},
	Timeout: 2 * time.Second,
})
```

### PowerShell

```go
result, err := sandbox.Run(context.Background(), sandbox.Config{
	Path:    "powershell",
	Args:    []string{"-Command", "echo 'Hello from PowerShell'"},
	Timeout: 2 * time.Second,
})
```

### Linux / macOS

```go
result, err := sandbox.Run(context.Background(), sandbox.Config{
	Path:    "sh",
	Args:    []string{"-c", "echo Hello from Unix"},
	Timeout: 2 * time.Second,
})
```

---

## Working Directory and Environment

```go
result, err := sandbox.Run(context.Background(), sandbox.Config{
	Path:    "go",
	Args:    []string{"env", "GOMOD"},
	Dir:     "/path/to/project",
	Env: []string{
		"GO111MODULE=on",
	},
	Timeout: 2 * time.Second,
})
```

This runs the command as if executed inside the specified directory
with the provided environment variables.
