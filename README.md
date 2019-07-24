## Hatewait [![Build Status](https://travis-ci.org/vkorobkov/hatewait.svg?branch=master)](https://travis-ci.org/vkorobkov/hatewait) [![Go Report Card](https://goreportcard.com/badge/github.com/vkorobkov/hatewait)](https://goreportcard.com/report/github.com/vkorobkov/hatewait)
Minimalistic "stopwatch" implementation. 
Written in Golang 1.12. Dependencies free. 

## Installation
```sh
go get github.com/vkorobkov/hatewait
```

## Examples

### Functional approach
Measure how many milliseconds does the job take:
```go
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
```
All available functions: 
```go
func DurationOf(x func()) time.Duration
func NanosecondsOf(x func()) int64
func MicrosecondsOf(x func()) int64
func MillisecondsOf(x func()) int64
func SecondsOf(x func()) int64
func HoursOf(x func()) int64
```

### Imperative approach
Measure how many milliseconds does the job take:
```go
import (
	"fmt"
	"github.com/vkorobkov/hatewait"
	"time"
)

func main() {
	startedAt := hatewait.Start()

	time.Sleep(1000 * 9000)

	elapsedMs := hatewait.MillisecondsSince(startedAt)

	fmt.Printf("The computation took %v milliseconds\n", elapsedMs)
}
```
All available functions:
```go
func DurationSince(since time.Time) time.Duration
func NanosecondsSince(since time.Time) int64
func MicrosecondsSince(since time.Time) int64
func MillisecondsSince(since time.Time) int64
func SecondsSince(since time.Time) int64
func HoursSince(since time.Time) int64
```
