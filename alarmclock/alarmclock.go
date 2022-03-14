package main

import (
	"fmt"
	"time"
)

func Remind(msg string, delay time.Duration) {
	for {
		time.Sleep(delay)
		fmt.Printf("The time is %s: %s\n", time.Now().Format("15.04.05"), msg)
	}
}

func main() {
	fmt.Println("Reminders started!")
	go Remind("Time to eat", time.Second*10)
	go Remind("Time to work", time.Second*30)
	Remind("Time to sleep", time.Second*60) // Last one runs on the main goroutine, blocks program from exiting
}
