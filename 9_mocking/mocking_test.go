package mocking

import (
	"bytes"
	"testing"
)

func TestCoundown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spysleeper := &SpySleeper{}

	Countdown(buffer, spysleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if spysleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3, got %d", spysleeper.Calls)
	}
}
