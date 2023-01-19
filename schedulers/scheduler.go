package schedulers

import (
	"github.com/noisyboy-9/colossus/process"
	"github.com/noisyboy-9/colossus/queue"
)

type Scheduler interface {
	Schedule(queue *queue.Queue) *process.Process
}
