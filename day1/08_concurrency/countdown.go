package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write ="write"
const sleep = "sleep"


type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func (s*SpySleeper) Sleep(){
	s.Calls++
	
}

type DefaultSleeper struct{}

func (d*DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}
 


func main()  {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i:=3; i>0; i--{
		fmt.Fprintln(out, i);
		sleeper.Sleep()
	}

	

	fmt.Fprint(out, "Go!")
}