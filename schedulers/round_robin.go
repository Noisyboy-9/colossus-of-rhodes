package schedulers

import (
	"time"

	"github.com/noisyboy-9/colossus-of-rhodes/preemptor/timers"
	"github.com/noisyboy-9/colossus-of-rhodes/process"
	"github.com/noisyboy-9/colossus-of-rhodes/queue"
)

const DefaultQuantum = 100 * time.Millisecond

type roundRobinScheduler struct {
	counter int
	quantum time.Duration
	timer   *timers.WatchDogTimer
}

func newRoundRobinScheduler(quantum time.Duration, timer *timers.WatchDogTimer) *roundRobinScheduler {
	return &roundRobinScheduler{
		quantum: quantum,
		timer:   timer,
	}
}

func (rs *roundRobinScheduler) GetType() SchedulerType {
	return ROUND_ROBIN
}

func (rs *roundRobinScheduler) Schedule(q queue.Queue) (selectedProcess *process.Process) {
	selectedProcess = q.GetByIndex(rs.counter)
	rs.counter = (rs.counter + 1) % queue.DegreeOfMultiprogramming

	rs.timer.Reset(rs.quantum)

	return selectedProcess
}
