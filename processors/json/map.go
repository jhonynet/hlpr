package json

import (
	"context"
	"fmt"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"sync"
)

var _ processor.Map = (*Processor)(nil)

func (r *Processor) CreateMap(p *pipeline.Pipeline, s *stages.Stage) processor.Map {
	return &Processor{
		pipeline: p,
		stage:    s,
	}
}

// todo: propagate context cancellation.
func (r *Processor) RunMap(_ context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(output)
		defer close(errChan)

		switch stage.Mode {
		case Decode:
			for data := range input {
				d, err := r.decode(data)
				if err != nil {
					errChan <- unit.Error{Err: err}
					continue
				}

				output <- d
			}
			return

		case Encode:
			for data := range input {
				d, err := r.encode(data)
				if err != nil {
					errChan <- unit.Error{Err: err}
					continue
				}

				output <- d
			}
			return
		}

		errChan <- unit.Error{Err: fmt.Errorf("%s is not a valid json mode", stage.Mode)}
	}()

	return output, errChan, nil
}
