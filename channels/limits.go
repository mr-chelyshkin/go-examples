package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // put empty to limit slot
	defer wg.Done()

	for j:=0;j<iterationsNum;j++ {
		fmt.Println(formatWork(in, j))

		// for even distribution of tasks
		//if j%2 == 0 {
		//	<-quotaCh              // return slot
		//	quotaCh <- struct{}{}  // get new slot
		//}

		runtime.Gosched()
	}
	<-quotaCh // read from limit slot
}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)

	for i:=0;i<goroutinesNum;i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func formatWork(in int, input int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "!!",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in, "recieved", input,
	)
}
