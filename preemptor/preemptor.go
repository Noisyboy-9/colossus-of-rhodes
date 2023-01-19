package preemptor

import "github.com/noisyboy-9/colossus/cpu"

type Preemptor interface {
	Preempte(cpu *cpu.CPU)
}
