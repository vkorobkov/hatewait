package main

import (
	"fmt"
	"github.com/vkorobkov/hatewait"
	"time"
)

func main() {
	elapsedMs := hatewait.MillisecondsOf(func() {
		time.Sleep(1000 * 9000)
	})

	fmt.Printf("The computation took %v milliseconds\n", elapsedMs)
}
