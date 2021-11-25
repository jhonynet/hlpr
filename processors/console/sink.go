package console

import (
	"context"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/unit"
	"github.com/jhonynet/hlpr/utils/logger"
)

var _ processor.Sink = (*Processor)(nil)

func (r *Processor) CreateSink(pipeline.Pipeline, pipeline.Stage) processor.Sink {
	return new(Processor)
}

func (r *Processor) RunSink(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan unit.Error, error) {
	errChan := make(chan unit.Error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(errChan)
		for entry := range input {
			logger.Info(ctx, spew.Sdump(entry))
		}
	}()

	return errChan, nil
}
