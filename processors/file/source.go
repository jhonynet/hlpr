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

var _ processor.Source = (*Processor)(nil)

func (r *Processor) CreateSource(p *pipeline.Pipeline, stage *stages.Stage) processor.Source {
	return &Processor{
		pipeline: p,
		stage:    stage,
	}
}

func (r *Processor) RunSource(ctx context.Context, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	f, err := os.Open(stage.Path)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot open %s file: %w", stage.Path, err)
	}

	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go r.read(ctx, stage, f, nil, output, errChan, wg)

	return output, errChan, nil
}
