// This is an example of handling system signals with
// channels

package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func workSignal(cancelChan <-chan os.Signal, notifyChan <-chan os.Signal) error {
	t := time.NewTicker(time.Second)

	for {
		select {
		case _ = <-cancelChan:
			t.Stop()

			return errors.New("User cancelled, exiting..")
		case _ = <-notifyChan:
			log.Printf("Noticed Ctrl-Z sent..")
		case <-t.C:
			log.Printf("Waiting for signal...")
		}
	}
}

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	cancelChan := make(chan os.Signal, 1)
	notifyChan := make(chan os.Signal, 1)

	signal.Notify(cancelChan, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(notifyChan, syscall.SIGTSTP)

	if err := workSignal(cancelChan, notifyChan); err != nil {
		log.Printf("Error: %s", err)
	}
}
