package main

import (
	"log"
	"sync"
	"time"
)

func producer() chan int {
	result := make(chan int)

	go func() {
		for i := 0; i < 16; i++ {
			result <- i
		}
		close(result)
	}()

	return result
}

func consumer(inputChan chan int) {
	for i := range inputChan {
		time.Sleep(time.Second)
		log.Printf("Processed item %d", i)
	}
}

func main() {
	inputChan := producer()

	log.Printf("Start working...")

	wg := &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			consumer(inputChan)
		}()
	}

	wg.Wait()
	log.Printf("All work done")
}
