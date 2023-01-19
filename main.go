package main

import (
	"fmt"
	"time"

	"github.com/noisyboy-9/colossus/dispatcher"
	"github.com/noisyboy-9/colossus/hardware"
	"github.com/noisyboy-9/colossus/process"
	"github.com/noisyboy-9/colossus/queue"
	"github.com/noisyboy-9/colossus/schedulers"
)

func main() {
	// setup hardware
	cpuReadyForNewTaskSignaller := make(chan bool)
	h := hardware.InitHardware(cpuReadyForNewTaskSignaller)

	// 	setup ready queue
	readyQueue := queue.NewReadyQueue()
	for i := 0; i < queue.DegreeOfMultiprogramming; i++ {
		p := process.NewProcess(int64(i), time.Now(), 500*time.Millisecond)
		readyQueue.Add(p)
	}

	// setup dispatcher
	dsp := dispatcher.NewSimpleDispatcher(h.Cpu)

	// 	setup scheduler
	scheduler := schedulers.NewScheduler(schedulers.ROUND_ROBIN, h)
	done := make(chan bool)
	go func() {
		for {
			switch {
			case <-done:
				fmt.Println("simulation finished")
				break

			case <-cpuReadyForNewTaskSignaller:
				selectedProcess := scheduler.Schedule(readyQueue)
				dsp.Dispatch(selectedProcess)
			}
		}
	}()

	// start simulation
	cpuReadyForNewTaskSignaller <- true
	time.Sleep(10 * 500 * time.Millisecond)
	done <- true
	fmt.Print("hope you have enjoyed")
}
