package main

import (
	"github.com/drewis/pb"
	"time"
)

func main() {
	count := int64(5000)
	bar := pb.New(count)

	// show percents (by default already true)
	bar.ShowPercent = true

	// show bar (by default already true)
	bar.ShowPercent = true

	// no need counters
	bar.ShowCounters = true

	bar.ShowTimeLeft = true

	// and start
	bar.Start()
	for i := int64(0); i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}
