package queues

import "github.com/noisyboy-9/colossus/process"

const DegreeOfMultiprogramming = 10

type Queue interface {
	Add(p *process.Process) error
	Remove(p *process.Process) error
	IsFull() bool
	IsEmpty() bool
}
