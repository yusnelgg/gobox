package sandbox

import "time"

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Duration time.Duration
	TimedOut bool
}
