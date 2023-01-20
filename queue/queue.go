package queue

import "github.com/noisyboy-9/colossus-of-rhodes/process"

const DegreeOfMultiprogramming = 10

type Queue interface {
	Add(p *process.Process) error
	Remove(p *process.Process) error
	IsFull() bool
	IsEmpty() bool
	GetByIndex(i int) *process.Process
}
