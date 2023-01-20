package timers

import "time"

type Timer interface {
	Start(d time.Duration)
	Stop()
	Reset(d time.Duration)
	Working() bool
	countTime()
}
