package main

import (
	"crypto/rand"
	"sync"
	"time"
)

const (
	duration   = 120
	process    = 12
	memoryInGb = 32
)

func main() {
	// Perform CPU intensive work

	go memoryInstesiveWork(process, duration, memoryInGb*1024)
	if err := cpuIntensive(process, duration); err != nil {
		// Handle error if needed
		panic(err)
	}
}

// cpuIntensive performs CPU intensive work for a given number of processes and duration
func cpuIntensive(process int, timeInSeconds int) error {
	exit := make(chan bool)
	go func() {
		// Sleep for the given time
		time.Sleep(time.Duration(timeInSeconds) * time.Second)
		close(exit)
	}()

	// Create a WaitGroup to wait for all the goroutines to finish
	wg := sync.WaitGroup{}

	for i := 0; i < process; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Perform CPU intensive work
			fibonacci(45, exit)
		}()
	}
	wg.Wait()
	return nil
}

// fibonacci calculates the Fibonacci number for a given n, with an exit channel to stop early
func fibonacci(n int, exit chan bool) int {
	select {
	case <-exit:
		return 0
	default:
		if n <= 1 {
			return n
		}
		return fibonacci(n-1, exit) + fibonacci(n-2, exit)
	}
}

func memoryInstesiveWork(workers int, timeInSeconds int, memoryInMb int) error {

	memoryperworker := memoryInMb / workers

	exit := make(chan bool)
	go func() {
		// Sleep for the given time
		time.Sleep(time.Duration(timeInSeconds) * time.Second)
		close(exit)
	}()

	// Create a WaitGroup to wait for all the goroutines to finish
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Perform CPU intensive work
			memoryLeak(memoryperworker, exit)
		}()
	}
	wg.Wait()
	return nil
}

func memoryLeak(memoryInMb int, exit chan bool) {

	select {
	case <-exit:
		return
	default:
		// Allocate memory
		buffer := make([]byte, memoryInMb*1024*1024)


		// Generate a random number
		rand.Read(buffer)


		//

	}
}
