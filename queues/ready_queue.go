package queues

import (
	"errors"
	"fmt"

	"github.com/noisyboy-9/colossus/process"
)

type ReadyQueue struct {
	length int
	queue  []*process.Process
}

func NewReadyQueue(length int) *ReadyQueue {
	return &ReadyQueue{
		length: length,
		queue:  make([]*process.Process, length),
	}
}

func (rq *ReadyQueue) AddProcess(p *process.Process) error {
	if len(rq.queue) >= rq.length {
		return errors.New("ready queue is full")
	}

	rq.queue = append(rq.queue, p)
	return nil
}

func (rq *ReadyQueue) RemoveProcess(target *process.Process) error {
	for i, p := range rq.queue {
		if target.Id() == p.Id() {
			rq.queue = append(rq.queue[:i], rq.queue[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("process with pid: %v doesn't exist", target.Id()))
}
