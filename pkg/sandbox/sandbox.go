package sandbox

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

func Run(ctx context.Context, cfg Config) (*Result, error) {
	start := time.Now()

	if cfg.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, cfg.Path, cfg.Args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Duration: time.Since(start),
	}

	if ctx.Err() == context.DeadlineExceeded {
		result.TimedOut = true
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			return result, err
		}
	}

	return result, nil
}
