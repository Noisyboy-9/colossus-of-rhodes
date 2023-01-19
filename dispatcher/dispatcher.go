package dispatcher

import (
	"github.com/noisyboy-9/colossus/cpu"
	"github.com/noisyboy-9/colossus/process"
)

type Dispatcher interface {
	Init(cpu *cpu.CPU)
	Dispatch(p *process.Process)
}
