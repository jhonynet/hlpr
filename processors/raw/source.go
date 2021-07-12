package raw

import (
	"context"
	"errors"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"sync"
)

var _ processor.Source = (*Processor)(nil)

func (r *Processor) CreateSource(_ *pipeline.Pipeline, stage *stages.Stage) processor.Source {
	return &Processor{
		stage: stage,
	}
}

func (r *Processor) RunSource(ctx context.Context, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage

	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	if len(stage.Data) == 0 {
		return nil, nil, errors.New("data is empty")
	}

	out := make(chan *unit.Data)
	errChan := make(chan unit.Error, 1)

	wg.Add(1)
	go func() {
		defer close(out)
		defer close(errChan)
		defer wg.Done()
		for _, line := range stage.Data {
			select {
			case out <- &unit.Data{Value: line}:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, errChan, nil
}
