package cpu

import (
	"errors"
	"time"

	"github.com/noisyboy-9/colossus-of-rhodes/process"
)

type CPU struct {
	executing          *process.Process
	interruptPin       chan bool
	ticker             *time.Ticker
	readyForNewProcess chan bool
}

func NewCPU(readySignaler chan bool) *CPU {
	return &CPU{
		interruptPin:       make(chan bool),
		ticker:             time.NewTicker(1 * time.Millisecond),
		readyForNewProcess: readySignaler,
	}
}

func (c *CPU) StartExecution(p *process.Process) {
	c.executing = p

	go func() {
		for {
			select {
			case <-c.interruptPin:
				c.preemptionRoutine()
				break
			case <-c.ticker.C:
				if c.executing.IsFinished() {
					_ = c.Interrupt()
				}
				c.executing.ReduceDuration(1 * time.Millisecond)
			}
		}
	}()
}

func (c *CPU) Interrupt() error {
	if c.executing == nil {
		return errors.New("no process is in execution")
	}

	c.interruptPin <- true
	return nil
}

func (c *CPU) preemptionRoutine() {
	c.ticker.Stop()
	c.executing = nil
	c.readyForNewProcess <- true
}
