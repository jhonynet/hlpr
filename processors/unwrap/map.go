package unwrap

import (
	"context"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"sync"
)

func (r *Processor) CreateMap(p *pipeline.Pipeline, s *stages.Stage) processor.Map {
	return &Processor{
		pipeline: p,
		stage:    s,
	}
}

// todo: propagate context cancellation.
func (r *Processor) RunMap(_ context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(output)
		defer close(errChan)

		for data := range input {
			if val, ok := data.Value.([]interface{}); ok {
				for _, v := range val {
					output <- &unit.Data{Value: v}
				}

				return
			}
			output <- data
		}
	}()

	return output, errChan, nil
}
