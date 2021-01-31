// This is as an example of receiving result from async functions
// using golang channel

package main

import (
	"log"
	"time"
)

// Result is the result type from the work functions
type Result struct {
	Val int
	Err error
}

// simulate very simple workload: sleep 1 second, return id + 1
func work(resultChan chan Result, id int) {
	time.Sleep(2 * time.Second)

	resultChan <- Result{Val: id + 1}
}

func main() {
	resultChan := make(chan Result)

	log.Print("work starting...")
	for i := 0; i < 5; i++ {
		// carefully: id here needs to be passed in as parameter
		// to async function
		// goroutines are started async, and may read variables
		// in closure after the iteration is finished and index i is updated
		go func(i int) {
			work(resultChan, i)
		}(i)
	}

	var results []int
	for i := 0; i < 5; i++ {
		res := <-resultChan
		if res.Err != nil {
			log.Printf("Error: unable to get result for %d: %s", i, res.Err)
			continue
		}

		results = append(results, res.Val)
	}
	log.Print("work finished")
	log.Printf("results: %v", results)
}
