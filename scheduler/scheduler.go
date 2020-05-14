package scheduler

import "time"

func Schedule(fun func(), sleepPeriod time.Duration) {
	for {
		fun()
		time.Sleep(sleepPeriod)
	}
}
