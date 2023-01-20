package timers

import "time"

type Timer interface {
	Start(d time.Duration)
	Stop()
	Reset()
	Working() bool
}
