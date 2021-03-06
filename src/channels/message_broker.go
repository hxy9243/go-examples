package main

import (
	"log"
	"sync"
	"time"
)

func worker(id int) {
	log.Printf("Starting work %d", id)
	time.Sleep(time.Second)
	log.Printf("Finishing work %d", id)
}

func main() {
	// create a slice of channels, that works as a message
	// broker from the main thread to worker goroutines
	chans := []chan struct{}{}
	log.Printf("Starting workers")

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		readyChan := make(chan struct{})
		chans = append(chans, readyChan)

		go func(id int, readyChan <-chan struct{}) {
			defer func() {
				wg.Done()
			}()

			<-readyChan
			worker(id)
		}(i, readyChan)
	}

	// message all the channels that the work is ready
	for _, ch := range chans {
		ch <- struct{}{}
	}

	wg.Wait()
	log.Printf("Finishing workers...")
}
