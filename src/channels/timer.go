package main

import (
	"fmt"
	"log"
	"time"
)

func workerTimer(t time.Duration, timeout time.Duration) error {
	ticker := time.NewTicker(t)
	timer := time.NewTimer(timeout)

	for {
		select {
		case _ = <-ticker.C:
			log.Printf("Still working...")
		case t := <-timer.C:
			ticker.Stop()

			return fmt.Errorf("Timeout at %s", t)
		}
	}
}

func main() {
	if err := workerTimer(time.Second, 5*time.Second); err != nil {
		log.Printf("Error: %s", err)
	}
}
