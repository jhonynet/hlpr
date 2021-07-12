package template

import (
	"bytes"
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

	// render template
	tpl, err := processor.NewTemplate(stage.Output, r.pipeline.Definition.Vars)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse template due an error %s", err)
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
			var buff bytes.Buffer
			if err := tpl.Render(&buff, data); err != nil {
				errChan <- unit.Error{Err: err}

				return
			}

			out <- data.SetValue(buff.String())
		}
	}()

	return out, errChan, nil
}
