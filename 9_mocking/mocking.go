package mocking

import (
	"fmt"
	"io"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) { // тут мы приняли интерфейс, который поможет тестировать
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second) // реальная отработка, но думаю мы заменим это на интерфейс
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
