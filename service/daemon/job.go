package daemon

import "time"

type Job struct {
	Duration time.Duration
	Func     func()
}
