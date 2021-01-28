package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperations int32

func inc() {
	atomic.AddInt32(&totalOperations, 1)
}

func main() {
	for i:=0;i<1000;i++ {
		go inc()
	}
	time.Sleep(2 *time.Millisecond)
	fmt.Println("total operations = ", totalOperations)
}

// atomic for simple operation like that
// faster than mutex (used inside mutex)
