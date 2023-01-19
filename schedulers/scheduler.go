package schedulers

import (
	"github.com/noisyboy-9/colossus/process"
	"github.com/noisyboy-9/colossus/queue"
)

type SchedulerType int

const (
	ROUND_ROBIN = iota
)

type Scheduler interface {
	GetType() SchedulerType
	Schedule(queue queue.Queue) *process.Process
}
