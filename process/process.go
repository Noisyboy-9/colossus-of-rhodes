package process

import "time"

type Process struct {
	id       int64
	start    time.Time
	duration time.Duration
}

func NewProcess(id int64, start time.Time, duration time.Duration) *Process {
	return &Process{
		id:       id,
		start:    start,
		duration: duration,
	}
}

func (p *Process) Id() int64 {
	return p.id
}

func (p *Process) Start() time.Time {
	return p.start
}

func (p *Process) Duration() time.Duration {
	return p.duration
}

func (p *Process) ReduceDuration(d time.Duration) {
	if d > p.duration {
		p.duration = 0
		return
	}

	p.duration -= d
}

func (p *Process) IsFinished() bool {
	return p.duration == 0
}
