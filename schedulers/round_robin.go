package schedulers

import (
	"time"

	"github.com/noisyboy-9/colossus/preemptor/timers"
	"github.com/noisyboy-9/colossus/process"
	"github.com/noisyboy-9/colossus/queue"
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

func (rs *roundRobinScheduler) Schedule(q queue.Queue) *process.Process {

	// stop previous watch dog timer
	p := q.GetByIndex(rs.counter)
	rs.counter = (rs.counter + 1) % queue.DegreeOfMultiprogramming

	err := rs.timer.StartTimer(rs.quantum)
	if err != nil {
		panic(err)
	}

	return p
}
