package template

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/unit"
)

var _ processor.Map = (*Processor)(nil)

func (r *Processor) CreateMap(pipeline pipeline.Pipeline, stage pipeline.Stage) processor.Map {
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

	// render template
	tpl, err := processor.NewTemplate(stage.Output, r.pipeline.Metadata.Globals)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse template due an error %s", err)
	}

	return processor.Mapper(ctx, input, r.mapFunc(tpl), wg)
}

func (r *Processor) mapFunc(tpl *processor.Template) processor.MapFunc {
	return func(ctx context.Context, data *unit.Data, output chan *unit.Data, errChan chan unit.Error) {
		var buff bytes.Buffer
		if err := tpl.Render(&buff, data); err != nil {
			errChan <- unit.Error{Err: err}

			return
		}

		output <- data.SetValue(buff.String())
	}
}
