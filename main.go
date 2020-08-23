package main

import (
	"fmt"
	"simple-workerserver-go/server"
	"time"
)

func main() {
	s := server.New()

	workToBeDone := &server.Work{
		Operation: func(a, b int) int { return a + b },
		A:         50,
		B:         10,
		Reply:     make(chan int),
	}

	s <- workToBeDone

	select {
	case res := <-workToBeDone.Reply:
		fmt.Println("The result of the work: ", res)
	case <-time.After(time.Second):
		fmt.Println("no result. Are you there?")
	}

}
