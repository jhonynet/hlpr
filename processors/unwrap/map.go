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
func (r *Processor) RunMap(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	return processor.Mapper(ctx, input, r.mapFunc(), wg)
}

func (r *Processor) mapFunc() processor.MapFunc {
	return func(_ context.Context, data *unit.Data, output chan *unit.Data, errChan chan unit.Error) {
		if val, ok := data.Value.([]interface{}); ok {
			for _, v := range val {
				output <- &unit.Data{Value: v}
			}

			return
		}
		output <- data
	}
}
