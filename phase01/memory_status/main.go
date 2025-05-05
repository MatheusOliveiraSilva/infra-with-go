package main

import (
	"os"
	"os/signal"
	"sync"
	"flag"
	"fmt"
	"time"
	"runtime"
)

func main() {

	var m runtime.MemStats // memory stats variable
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
				runtime.ReadMemStats(&m)
				fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
				fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
				fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
				fmt.Printf("\tNumGC = %v\n", m.NumGC)
			case <-signalChannel:
				fmt.Println("Received interrupt signal, shutting down...")
				return
			}
		}
	}()

	fmt.Printf("Getting memory stats each %d seconds...\n", *interval)
	wg.Wait()
	fmt.Println("Goroutine finished")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}