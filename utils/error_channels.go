package utils

import (
	"sync"

	"github.com/jhonynet/hlpr/unit"
)

func MergeErrors(errorChannels ...<-chan unit.Error) <-chan unit.Error {
	var (
		wg  sync.WaitGroup
		out = make(chan unit.Error, len(errorChannels))
	)

	for _, c := range errorChannels {
		wg.Add(1)
		go func(errChan <-chan unit.Error) {
			for n := range errChan {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
