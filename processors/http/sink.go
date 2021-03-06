package http

import (
	"context"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"sync"
)

var _ processor.Sink = (*Processor)(nil)

func (r *Processor) CreateSink(pipeline *pipeline.Pipeline, stage *stages.Stage) processor.Sink {
	return &Processor{
		stage:    stage,
		pipeline: pipeline,
	}
}

func (r *Processor) RunSink(context.Context, <-chan *unit.Data, *sync.WaitGroup) (<-chan unit.Error, error) {
	panic("implement me")
}
