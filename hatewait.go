package hatewait

import "time"

func Start() time.Time {
	return time.Now()
}

func DurationSince(since time.Time) time.Duration {
	return time.Since(since)
}

func NanosecondsSince(since time.Time) int64 {
	return DurationSince(since).Nanoseconds()
}

func MicrosecondsSince(since time.Time) int64 {
	return NanosecondsSince(since) / int64(time.Microsecond)
}

func MillisecondsSince(since time.Time) int64 {
	return NanosecondsSince(since) / int64(time.Millisecond)
}

func SecondsSince(since time.Time) int64 {
	return NanosecondsSince(since) / int64(time.Second)
}

func HoursSince(since time.Time) int64 {
	return NanosecondsSince(since) / int64(time.Hour)
}

func DurationOf(x func()) time.Duration {
	nanoseconds := measure(x, NanosecondsSince)
	return time.Duration(nanoseconds)
}

func NanosecondsOf(x func()) int64 {
	return measure(x, NanosecondsSince)
}

func MicrosecondsOf(x func()) int64 {
	return measure(x, MicrosecondsSince)
}

func MillisecondsOf(x func()) int64 {
	return measure(x, MillisecondsSince)
}

func SecondsOf(x func()) int64 {
	return measure(x, SecondsSince)
}

func HoursOf(x func()) int64 {
	return measure(x, HoursSince)
}

func measure(x func(), converter func(time.Time) int64) int64 {
	startedAt := Start()
	x()
	return converter(startedAt)
}
