package sandbox

import "time"

type Config struct {
	Path    string
	Args    []string
	Timeout time.Duration
	Dir     string
	Env     []string
}
