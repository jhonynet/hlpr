package file

import (
	"context"
	"fmt"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"os"
	"sync"
)

var _ processor.Sink = (*Processor)(nil)

func (r *Processor) CreateSink(p *pipeline.Pipeline, stage *stages.Stage) processor.Sink {
	return &Processor{
		pipeline: p,
		stage:    stage,
	}
}

func (r *Processor) RunSink(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, err
	}

	f, err := os.Create(stage.Path)
	if err != nil {
		return nil, fmt.Errorf("cannot create file %s", err)
	}

	errChan := make(chan unit.Error, 1)
	wg.Add(1)
	go r.write(ctx, stage, f, input, nil, errChan, wg)

	return errChan, nil
}
