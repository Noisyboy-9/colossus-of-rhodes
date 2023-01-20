package timers

import (
	"time"

	"github.com/noisyboy-9/colossus-of-rhodes/cpu"
)

type WatchDogTimer struct {
	ticker  *time.Ticker
	done    chan bool
	working bool
	cpu     *cpu.CPU
}

func NewWatchDogTimer(cpu *cpu.CPU) *WatchDogTimer {
	return &WatchDogTimer{
		cpu:     cpu,
		done:    make(chan bool),
		working: false,
		ticker:  nil,
	}
}

func (w *WatchDogTimer) countTime() {
	for {
		select {
		case <-w.ticker.C:
			w.Preempte()
		case <-w.done:
			w.working = false
			w.ticker.Stop()
			break
		}
	}
}

func (w *WatchDogTimer) Preempte() {
	err := w.cpu.Interrupt()
	if err != nil {
		panic(err)
	}
}

func (w *WatchDogTimer) Start(d time.Duration) {
	if w.Working() {
		panic("start on already started watchdog timer")
	}
	if w.ticker != nil {
		w.ticker = time.NewTicker(d)
	}
	go w.countTime()
}

func (w *WatchDogTimer) Stop() {
	if !w.Working() {
		panic("stop on already stopped watchdog timer")
	}
	w.done <- true
}

func (w *WatchDogTimer) Reset(d time.Duration) {
	if w.ticker == nil {
		w.ticker = time.NewTicker(d)
	} else {
		w.ticker.Reset(d)
	}

	go w.countTime()
}

func (w *WatchDogTimer) Working() bool {
	return w.working
}
