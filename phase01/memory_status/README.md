# Memory Status

A cross-platform utility that monitors and displays memory statistics of the Go runtime environment.

## Features

- Uses Go's runtime package to collect memory statistics
- Displays memory allocation, total allocation, system memory, and garbage collection metrics
- Runs on any platform supported by Go
- Configurable polling interval via command-line flags
- Handles OS interrupt signals for graceful shutdown

## Usage

Build the program:
```
go build -o memstat
```

Run with default interval (1 second):
```
./memstat
```

Specify a custom interval:
```
./memstat -interval=5
```

## Example Output

```
Getting memory stats each 1 seconds...
Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 8 MiB	NumGC = 0
Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 8 MiB	NumGC = 0
Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 8 MiB	NumGC = 0
^C
Received interrupt signal, shutting down...
Goroutine finished
```

## Memory Stat Fields Explained

- **Alloc**: Memory currently allocated and not yet freed (in MiB)
- **TotalAlloc**: Total memory allocated since program start, even if freed (in MiB)
- **Sys**: Memory obtained from the OS (in MiB)
- **NumGC**: Number of completed garbage collection cycles 