// This is a toy demo of implementing "future" like interface
// with golang channel primitives

package main

import (
	"fmt"
	"log"
	"time"
)

// Future is a wrapper for channel, to return future result in a channel
type Future struct {
	Chan chan interface{}
	Err  error
}

// NewFuture returns a new Future struct with channel initialized
func NewFuture() Future {
	return Future{
		Chan: make(chan interface{}),
	}
}

// Get waits and gets the result of future. Return error if there's error in the result
func (fut *Future) Get() (interface{}, error) {
	if fut.Err != nil {
		return nil, fut.Err
	}

	result := <-fut.Chan
	close(fut.Chan)

	return result, nil
}

// work is an example of using Future as function return
func work(i int) Future {
	result := NewFuture()
	go func() {
		time.Sleep(2 * time.Second)

		result.Chan <- fmt.Sprintf("Hello World %d", i)
	}()
	return result
}

// this example demos a very simple "future"-like use case of golang
// channels, by returning a future from calling async procedures, then wait
// and collect all the results
func main() {
	futs := []Future{}

	log.Printf("Starting work async...")
	for i := 0; i < 5; i++ {
		futs = append(futs, work(i))
	}

	for i, fut := range futs {
		if val, err := fut.Get(); err != nil {
			log.Printf("ERRO: getting result %d: %s", i, err)
		} else {
			result := val.(string)

			log.Printf("INFO: Getting result %d: %s", i, result)
		}
	}
	log.Print("All work finished...")
}
