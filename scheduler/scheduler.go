package main

import (
	"fmt"
	"time"
)

func checkTime() int {
	currentTime := time.Now()
	time := currentTime.Hour()
	return time
}

func main() {
	var currentTime, startWindow, endWindow int
	currentTime = checkTime()
	startWindow = 18
	endWindow = 23

	if currentTime >= startWindow && currentTime < endWindow {
		fmt.Println("Within window")
		StartingInstance()
	} else {
		StoppingInstance()
	}

}
