package http

import (
	"context"
	"fmt"
	"sync"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/unit"
)

var _ processor.Source = (*Processor)(nil)

func (r *Processor) CreateSource(pipeline pipeline.Pipeline, stage pipeline.Stage) processor.Source {
	return &Processor{
		stage:    stage,
		pipeline: pipeline,
	}
}

func (r *Processor) RunSource(ctx context.Context, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	// parse url template
	urlTpl, err := processor.NewTemplate(stage.URL, r.pipeline.Metadata.Globals)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse http URL %s %w", stage.URL, err)
	}

	// parse body template if any
	bodyTpl, err := processor.NewTemplate(stage.Body, r.pipeline.Metadata.Globals)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse http body %w", err)
	}

	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(output)
		defer close(errChan)

		res, err := r.processUnit(ctx, stage, urlTpl, bodyTpl, &unit.Data{})
		if err != nil {
			errChan <- unit.Error{Err: err}
			return
		}

		output <- res
	}()

	return output, errChan, nil
}
