package processor

import (
	"context"
	"github.com/jhonynet/hlpr/unit"
	"sync"
)

type MapFunc func(context.Context, *unit.Data, chan *unit.Data, chan unit.Error)

// todo: propagate context cancellation.
func Mapper(ctx context.Context, input <-chan *unit.Data, cb MapFunc, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(output)
		defer close(errChan)

		for data := range input {
			cb(ctx, data, output, errChan)
		}
	}()

	return output, errChan, nil
}
