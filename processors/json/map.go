package json

import (
	"context"
	"sync"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/unit"
)

var _ processor.Map = (*Processor)(nil)

func (r *Processor) CreateMap(p pipeline.Pipeline, s pipeline.Stage) processor.Map {
	return &Processor{
		pipeline: p,
		stage:    s,
	}
}

// todo: propagate context cancellation.
func (r *Processor) RunMap(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	return processor.Mapper(ctx, input, r.mapFunc(stage), wg)
}

func (r *Processor) mapFunc(stage Stage) processor.MapFunc {
	return func(_ context.Context, data *unit.Data, output chan *unit.Data, errChan chan unit.Error) {
		fn := r.decode
		if stage.Mode == Encode {
			fn = r.encode
		}

		d, err := fn(data)

		if err != nil {
			errChan <- unit.Error{Err: err}
			return
		}

		output <- d
	}
}
