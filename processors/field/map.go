package field

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

func (r *Processor) RunMap(_ context.Context, input <-chan *unit.Data, wg *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error) {
	var stage Stage
	if err := r.stage.Bind(&stage); err != nil {
		return nil, nil, err
	}

	if stage.Name == "" {
		return nil, nil, fmt.Errorf("name property is required")
	}

	if len(stage.Stages) == 0 {
		return nil, nil, fmt.Errorf("stage property is required")
	}

	processorProvider = new(processor.Provider)
	for idx, stage := range stage.Stages {

	}

	var (
		out     = make(chan *unit.Data)
		errChan = make(chan unit.Error, 1)
	)

	wg.Add(1)
	go func() {
		defer close(errChan)
		defer close(out)
		defer wg.Done()

		for data := range input {
			// todo: limit input to be a map

			// create a channel to get the value


			// set the stage.Name property to data
			d, err := data.SetProperty(stage.Name, "value")
			if err != nil {
				errChan <- unit.Error{Err: err}

				return
			}

			// send data to output channel
			out <- d
		}
	}()

	return out, errChan, nil
}
