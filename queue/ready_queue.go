package queue

import (
	"errors"
	"fmt"

	"github.com/noisyboy-9/colossus-of-rhodes/process"
)

type ReadyQueue struct {
	processes []*process.Process
	capacity  int
}

func NewReadyQueue() *ReadyQueue {
	return &ReadyQueue{
		processes: make([]*process.Process, 0),
		capacity:  DegreeOfMultiprogramming,
	}
}
func (r *ReadyQueue) GetByIndex(index int) *process.Process {
	return r.processes[index]
}

func (r *ReadyQueue) Add(p *process.Process) error {
	if r.IsFull() {
		return errors.New(fmt.Sprintf("ready queue with degree of multiprogramming: %v is full", DegreeOfMultiprogramming))
	}

	r.processes = append(r.processes, p)
	return nil
}

func (r *ReadyQueue) Remove(target *process.Process) error {
	if r.IsEmpty() {
		return errors.New("call remove process on empty queue")
	}

	for i, p := range r.processes {
		if p.Id() == target.Id() {
			r.processes = append(r.processes[:i], r.processes[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("error with pid: %v can't be found in ready queue", target.Id()))
}

func (r *ReadyQueue) IsFull() bool {
	return len(r.processes) == r.capacity
}

func (r *ReadyQueue) IsEmpty() bool {
	return len(r.processes) == 0
}
