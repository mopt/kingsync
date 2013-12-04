package syncers

import (
	"time"
)

type SyncerNop struct {
}

func (s *SyncerNop) Sync(files []string, target string, progress chan int) {
	go func() {
		for i := 0; i <= 100; i += 5 {
			time.Sleep(20 * time.Millisecond)
			progress <- i
		}
	}()
}
