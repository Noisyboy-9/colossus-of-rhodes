package hardware

import (
	"github.com/noisyboy-9/colossus-of-rhodes/cpu"
	"github.com/noisyboy-9/colossus-of-rhodes/preemptor/timers"
)

type Hardware struct {
	Cpu           *cpu.CPU
	WatchdogTimer *timers.WatchDogTimer
}

func InitHardware(readySignaller chan bool) *Hardware {
	c := cpu.NewCPU(readySignaller)
	return &Hardware{
		Cpu:           c,
		WatchdogTimer: timers.NewWatchDogTimer(c),
	}
}
