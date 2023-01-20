package schedulers

import (
	"fmt"

	"github.com/noisyboy-9/colossus-of-rhodes/hardware"
	"github.com/noisyboy-9/colossus-of-rhodes/process"
	"github.com/noisyboy-9/colossus-of-rhodes/queue"
)

type SchedulerType int

const (
	ROUND_ROBIN = iota
)

type Scheduler interface {
	GetType() SchedulerType
	Schedule(queue queue.Queue) *process.Process
}

func NewScheduler(t SchedulerType, hardware *hardware.Hardware) Scheduler {
	switch t {
	case ROUND_ROBIN:
		return newRoundRobinScheduler(DefaultQuantum, hardware.WatchdogTimer)
	default:
		panic(fmt.Sprintf("unknown scheduler type: %d", t))
	}
}
