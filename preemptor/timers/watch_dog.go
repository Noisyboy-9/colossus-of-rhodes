package timers

import (
	"errors"
	"fmt"
	"time"

	"github.com/noisyboy-9/colossus/cpu"
)

type WatchDogTimer struct {
	ticker      *time.Ticker
	cpu         *cpu.CPU
	stopChannel chan bool
}

func (t *WatchDogTimer) stopRoutine() error {
	if !t.IsWorking() {
		return errors.New("stop routine called on stopped timer")
	}

	t.ticker.Stop()
	t.ticker = nil
	return nil
}

func NewWatchDogTimer(c *cpu.CPU) *WatchDogTimer {
	return &WatchDogTimer{
		ticker:      nil,
		cpu:         c,
		stopChannel: make(chan bool),
	}
}

func (t *WatchDogTimer) StartTimer(timeToGoOff time.Duration) error {
	if t.IsWorking() {
		return errors.New("start timer call on already started timer")
	}

	t.ticker = time.NewTicker(timeToGoOff)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.Preempte()
			case <-t.stopChannel:
				err := t.stopRoutine()
				if err != nil {
					panic(err)
				}
				break
			}
		}
	}()

	return nil
}

func (t *WatchDogTimer) IsWorking() bool {
	return t.ticker != nil
}

func (t *WatchDogTimer) StopTimer() error {
	if !t.IsWorking() {
		return errors.New("stop timer called on empty timer")
	}
	t.stopChannel <- true
	t.ticker = nil
	return nil
}

func (t *WatchDogTimer) Preempte() {
	err := t.cpu.Interrupt()
	if err != nil {
		panic(fmt.Sprintf("watchdog blew on empty cpu: %v", err.Error()))
	}

	err = t.StopTimer()
	if err != nil {
		panic(err)
	}
}
