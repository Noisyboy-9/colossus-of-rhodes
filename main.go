package main

import (
	"fmt"
	"time"

	"github.com/noisyboy-9/colossus-of-rhodes/dispatcher"
	"github.com/noisyboy-9/colossus-of-rhodes/hardware"
	"github.com/noisyboy-9/colossus-of-rhodes/process"
	"github.com/noisyboy-9/colossus-of-rhodes/queue"
	"github.com/noisyboy-9/colossus-of-rhodes/schedulers"
)

func main() {
	// setup hardware
	cpuReadyForNewTaskSignaller := make(chan bool, 1)
	h := hardware.InitHardware(cpuReadyForNewTaskSignaller)

	// 	setup ready queue
	readyQueue := queue.NewReadyQueue()
	for i := 0; i < queue.DegreeOfMultiprogramming; i++ {
		p := process.NewProcess(int64(i), time.Now(), 500*time.Millisecond)
		err := readyQueue.Add(p)
		if err != nil {
			panic(err)
		}
	}

	// setup dispatcher
	dsp := dispatcher.NewSimpleDispatcher(h.Cpu)

	// 	setup scheduler
	scheduler := schedulers.NewScheduler(schedulers.ROUND_ROBIN, h)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("simulation finished")
				break

			case <-cpuReadyForNewTaskSignaller:
				selectedProcess := scheduler.Schedule(readyQueue)
				fmt.Printf("process with pid: %v selected\n", selectedProcess.Id())
				dsp.Dispatch(selectedProcess)
			}
		}
	}()

	// start simulation
	cpuReadyForNewTaskSignaller <- true
	time.Sleep(queue.DegreeOfMultiprogramming * 500 * time.Millisecond)
	done <- true
	fmt.Println("hope you have enjoyed")
}
