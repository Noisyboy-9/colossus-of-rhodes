package storages

import "github.com/noisyboy-9/colossus/process"

type Queue interface {
	Add(p *process.Process) error
	Remove(p *process.Process) error
	IsFull() bool
	IsEmpty() bool
}
