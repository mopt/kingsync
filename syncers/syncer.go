package syncers

// A Syncer represents a backend that is capable of synchronizing files with
// the given target. It should return immediately; starting goroutines in the
// background to handle synchronization. Progress must be reported through
// the provided channel in the form of an integer from 0 to 100.
type Syncer interface {
	Sync(files []string, target string, progress chan int)
}

// Desired implementations:
//  - Nop
//  - Hardlink
//  - Rsync
