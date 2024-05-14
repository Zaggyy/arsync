package main

import "time"

var sleepTime = 5 * time.Second

func setSleepTime(time time.Duration) {
	sleepTime = time
}

func sleep() {
	time.Sleep(sleepTime)
}
