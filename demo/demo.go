package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	for i:=1; i<10;i++{
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main(){
	go printNumbers()
	fmt.Println("Main function is finished")
	time.Sleep(time.Second * 5)
}