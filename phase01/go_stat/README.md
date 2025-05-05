# Go Stat

A simple CLI utility that demonstrates basic Go concurrency patterns by printing "tick" at specified intervals.

## Features

- Uses goroutines for concurrent execution
- Implements channel-based communication
- Handles OS interrupt signals for graceful shutdown
- Configurable via command-line flags

## Usage

Build the program:
```
go build -o gostat
```

Run with default interval (1 second):
```
./gostat
```

Specify a custom interval:
```
./gostat -interval=2
```

## Example Output

```
Waiting for goroutine to finish...
tick
tick
tick
^CReceived interrupt signal, shutting down...
Goroutine finished
``` 