package hatewait

import (
	"testing"
)

func TestStart(t *testing.T) {
	Start() // at least does not fail
}

func TestNanosecondsSince(t *testing.T) {
	if NanosecondsSince(Start()) < 0 {
		t.Error("Must be zero or a positive")
	}
}

func TestMicrosecondsSince(t *testing.T) {
	if MicrosecondsSince(Start()) < 0 {
		t.Error("Must be zero or a positive")
	}
}

func TestMillisecondsSince(t *testing.T) {
	if MillisecondsSince(Start()) < 0 {
		t.Error("Must be zero or a positive")
	}
}

func TestSecondsSince(t *testing.T) {
	if SecondsSince(Start()) < 0 {
		t.Error("Must be zero or a positive")
	}
}

func TestHoursSince(t *testing.T) {
	if HoursSince(Start()) < 0 {
		t.Error("Must be zero or a positive")
	}
}

func TestDurationOf(t *testing.T) {
	callbackInvoked := false
	if DurationOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}

func TestNanosecondsOf(t *testing.T) {
	callbackInvoked := false
	if NanosecondsOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}

func TestMicrosecondsOf(t *testing.T) {
	callbackInvoked := false
	if MicrosecondsOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}

func TestMillisecondsOf(t *testing.T) {
	callbackInvoked := false
	if MillisecondsOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}

func TestSecondsOf(t *testing.T) {
	callbackInvoked := false
	if SecondsOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}

func TestHoursOf(t *testing.T) {
	callbackInvoked := false
	if HoursOf(func() { callbackInvoked = true }) < 0 {
		t.Error("Must be zero or a positive")
	}
	if !callbackInvoked {
		t.Error("Callback must be called")
	}
}
