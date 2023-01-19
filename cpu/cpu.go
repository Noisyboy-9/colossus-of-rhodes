package cpu

import (
	"github.com/noisyboy-9/colossus/process"
)

type CPU struct {
	executing *process.Process
}

func NewCPU() *CPU {
	return &CPU{}
}

func (c *CPU) StartExecution(p *process.Process) {
	c.executing = p
}
