package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func workload(ctx context.Context) error {
	t := time.NewTicker(3 * time.Second)
	log.Printf("Start work at %v", time.Now())

	for {
		select {
		case now := <-t.C:
			log.Printf("Still working at %v", now)
		case <-ctx.Done():
			log.Printf("Received cancel signal, exiting...")

			// when handling cancel signals, remember to properly close resources
			// to avoid resource leaks

			return errors.New("Context cancelled, stopping..")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 7*time.Second)
	defer cancel()

	if err := workload(ctx); err != nil {
		log.Printf("Error: %s", err)
	}
}
