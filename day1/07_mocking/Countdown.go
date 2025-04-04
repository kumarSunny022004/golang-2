package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Constants for Countdown
const countdownStart = 3
const finalWord = "Go!"

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// DefaultSleeper uses real time.Sleep()
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown function
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// Main function
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
