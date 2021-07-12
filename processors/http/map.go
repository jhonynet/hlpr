package http

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

func (r *Processor) CreateMap(pipeline *pipeline.Pipeline, stage *stages.Stage) processor.Map {
	return &Processor{
		stage:    stage,
		pipeline: pipeline,
	}
}

func (r *Processor) RunMap(ctx context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	// parse url template
	urlTpl, err := processor.NewTemplate(stage.URL, r.pipeline.Definition.Vars)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse http URL %s %w", stage.URL, err)
	}

	// parse body template if any
	bodyTpl, err := processor.NewTemplate(stage.Body, r.pipeline.Definition.Vars)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse http body %s", err)
	}

	errChan := make(chan unit.Error, 1)
	output := make(chan *unit.Data)
	wg.Add(1)
	go r.runWorkers(ctx, stage, urlTpl, bodyTpl, input, output, errChan, wg)

	return output, errChan, nil
}
