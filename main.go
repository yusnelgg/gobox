package main

import (
	"context"
	"time"

	"github.com/yusnelgg/gobox/pkg/sandbox"
)

func RunSandboxedCommand(path string, args []string, timeoutSeconds int) (*sandbox.Result, error) {
	cfg := sandbox.Config{
		Path:    path,
		Args:    args,
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
	return sandbox.Run(context.Background(), cfg)
}

func main() {

	// example usage
	result, err := sandbox.Run(context.Background(), sandbox.Config{
		Path: "go",
		Args: []string{"env", "GOMOD"},
		Dir:  "C:/Users/ExampleUser/Workspace",
		Env: []string{
			"GO111MODULE=on",
		},
		Timeout: 2 * time.Second,
	})
	if err != nil {
		println("Error executing command:", err.Error())
		return
	}
	println("Stdout:", result.Stdout)
	println("Stderr:", result.Stderr)
	println("Exit Code:", result.ExitCode)
	println("Duration (ms):", result.Duration.Milliseconds())
	println("Timed Out:", result.TimedOut)
}
