package core

import "time"

type Timer struct {
	Interval int
	Cycle    int
	Tick     func()
	running  bool
}

var timers = make(map[string]*Timer)

func SetTimer(name string, timer *Timer) string {
	if current, exists := timers[name]; exists {
		current.running = false
	}
	timers[name] = timer
	timer.running = true
	// JUST TEST CODE, THIS WILL BE EXECUTED BY THE GAME LOOP
	go func() {
		t := time.NewTicker(time.Second)
		for timer.running {
			timer.Tick()
			<-t.C
		}
	}()
	return name
}

func StopTimer(name string) {
	if timer, ok := timers[name]; ok {
		timer.running = false
	}
}