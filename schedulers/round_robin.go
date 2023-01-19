package schedulers

import (
	"time"

	"github.com/noisyboy-9/colossus/preemptor/timers"
	"github.com/noisyboy-9/colossus/process"
	"github.com/noisyboy-9/colossus/queue"
)

const DefaultQuantum = 100 * time.Millisecond

type RoundRobinScheduler struct {
	counter int
	quantum time.Duration
	timer   *timers.WatchDogTimer
}

func NewRoundRobinScheduler(quantum time.Duration, timer *timers.WatchDogTimer) *RoundRobinScheduler {
	return &RoundRobinScheduler{
		quantum: quantum,
		timer:   timer,
	}
}

func (rs *RoundRobinScheduler) GetType() SchedulerType {
	return ROUND_ROBIN
}

func (rs *RoundRobinScheduler) Schedule(q queue.Queue) *process.Process {
	p := q.GetByIndex(rs.counter)
	rs.counter = (rs.counter + 1) % queue.DegreeOfMultiprogramming

	err := rs.timer.StartTimer(rs.quantum)
	if err != nil {
		panic(err)
	}

	return p
}
