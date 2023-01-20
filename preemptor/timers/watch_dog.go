package timers

import (
	"time"
)

type WatchDogTimer struct {
	ticker  *time.Ticker
	done    chan bool
	working bool
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

}

func (w *WatchDogTimer) Start(d time.Duration) {
	w.ticker = time.NewTicker(d)

	go w.countTime()
}

func (w *WatchDogTimer) Stop() {
	w.done <- true
}

func (w *WatchDogTimer) Reset(d time.Duration) {
	w.ticker.Reset(d)

	go w.countTime()
}

func (w *WatchDogTimer) Working() bool {
	return w.working
}
