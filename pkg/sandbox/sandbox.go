package sandbox

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"time"
)

func Run(ctx context.Context, cfg Config) (*Result, error) {
	start := time.Now()

	if cfg.Path == "" {
		return nil, errors.New("sandbox: empty path")
	}

	if cfg.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, cfg.Path, cfg.Args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if cfg.Dir != "" {
		cmd.Dir = cfg.Dir
	}

	if len(cfg.Env) > 0 {
		cmd.Env = cfg.Env
	}

	err := cmd.Run()

	result := &Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Duration: time.Since(start),
		ExitCode: 0,
	}

	if ctx.Err() == context.DeadlineExceeded {
		result.TimedOut = true
		result.ExitCode = -1
		return result, nil
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
			return result, nil
		}
		return result, err
	}

	return result, nil
}
