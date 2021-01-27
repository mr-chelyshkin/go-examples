package main

import (
	"fmt"
	"time"
)

func LongQuery() <-chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(2 *time.Second)
		ch <- true
	}()
	return ch
}

func main() {
	timer := time.NewTimer(3 *time.Second)

	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happened")
	case result := <-LongQuery():
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result: ", result)
	}
}

