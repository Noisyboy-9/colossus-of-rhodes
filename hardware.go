package main

import (
	"github.com/noisyboy-9/colossus/cpu"
	"github.com/noisyboy-9/colossus/preemptor/timers"
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
