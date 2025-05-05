package main

import (
	"os"
	"os/signal"
	"sync"
	"flag"
	"fmt"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	interval := flag.Int("interval", 1, "interval in seconds")
	flag.Parse()

	channel := time.Tick(time.Second * time.Duration(*interval))

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-channel:
				fmt.Println("tick")
			case <-signalChannel:
				fmt.Println("Received interrupt signal, shutting down...")
				return
			}
		}
	}()

	fmt.Println("Waiting for goroutine to finish...")
	wg.Wait()
	fmt.Println("Goroutine finished")
}
