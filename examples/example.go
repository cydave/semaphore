package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cydave/semaphore"
)

func doWork(workerNum int, sem semaphore.Semaphore) {
	defer sem.Release()
	fmt.Printf("Worker %d processing\n", workerNum)
	n := rand.Intn(500)
	time.Sleep(time.Duration(n * int(time.Millisecond)))
	time.Sleep(time.Duration(n * int(time.Millisecond)))
	fmt.Printf("Worker %d finished\n", workerNum)
}

func main() {
	nWorkers := 10
	sem := semaphore.New(nWorkers)
	defer sem.Close()
	for i := 0; i < 100; i++ {
		sem.Acquire()
		go doWork((i%nWorkers)+1, sem)
	}
	fmt.Println("Done!")
}
