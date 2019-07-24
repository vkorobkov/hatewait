package main

import (
	"fmt"
	"github.com/vkorobkov/hatewait"
	"time"
)

func main() {
	functional()
	imperative()
}

func functional() {
	elapsedMs := hatewait.MillisecondsOf(func() {
		time.Sleep(1000 * 9000)
	})

	fmt.Printf("The computation took %v milliseconds\n", elapsedMs)
}

func imperative() {
	startedAt := hatewait.Start()

	time.Sleep(1000 * 9000)

	elapsedMs := hatewait.MillisecondsSince(startedAt)

	fmt.Printf("The computation took %v milliseconds\n", elapsedMs)
}
