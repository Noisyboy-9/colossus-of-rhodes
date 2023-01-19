package dispatcher

import (
	"github.com/noisyboy-9/colossus/cpu"
	"github.com/noisyboy-9/colossus/process"
)

type SimpleDispatcher struct {
	target *cpu.CPU
}

func NewSimpleDispatcher(target *cpu.CPU) *SimpleDispatcher {
	return &SimpleDispatcher{target: target}

}

func (s *SimpleDispatcher) Dispatch(p *process.Process) {
	s.target.StartExecution(p)
}
