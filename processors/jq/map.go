package jq

import (
	"context"
	"fmt"
	"sync"

	"github.com/itchyny/gojq"
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

func (r *Processor) RunMap(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	query, err := gojq.Parse(stage.Expr)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse jq query: %w", err)
	}

	return processor.Mapper(ctx, input, r.mapFunc(query), wg)
}

func (r *Processor) mapFunc(query *gojq.Query) processor.MapFunc {
	return func(ctx context.Context, data *unit.Data, output chan *unit.Data, errChan chan unit.Error) {
		iter := query.RunWithContext(ctx, data.Value)
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			if err, ok := v.(error); ok {
				errChan <- unit.ErrorFrom(fmt.Errorf("error happened during jq iteration %w", err))
				continue
			}

			output <- &unit.Data{Value: v}
		}
	}
}
