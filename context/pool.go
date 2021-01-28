package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func printFinishWork(workerNum int) {
	fmt.Sprintln("--- finish worker ", workerNum)
}

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Println(formatWork(workerNum, input))
		runtime.Gosched()
	}
	printFinishWork(workerNum)
}

func main() {
	workerInput := make(chan string, 2)
	for i:=0;i<goroutinesNum;i++ {
		go startWorker(i, workerInput)
	}

	items := []string{"one", "two", "three", "four", "five", "fix", "seven", "eight", "nine", "ten"}

	for _, item := range items {
		workerInput <-item
	}
	close(workerInput)
	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in, "recieved", input,
	)
}