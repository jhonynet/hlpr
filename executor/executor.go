package executor

import (
	"context"
	"fmt"
	"sync"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/utils/logger"
	"github.com/jhonynet/hlpr/workload"
	"go.uber.org/zap"
)

type defaultExecutor struct {
	processorRegistry processor.Registry
}

func NewDefaultExecutor(registry processor.Registry) Executor {
	return &defaultExecutor{
		processorRegistry: registry,
	}
}

func (d defaultExecutor) Execute(ctx context.Context, pipeline pipeline.Pipeline) error {
	processorProvider := new(processor.Provider)

	// configure each processor
	for idx, stage := range pipeline.Stages {
		switch true {

		case idx == 0: //first is source
			proc := d.processorRegistry.Get(stage)
			if proc == nil {
				return fmt.Errorf("processor for stage %s not found", stage.Type())
			}

			if sourceProc, ok := proc.(processor.Source); ok {
				processorProvider.Source = sourceProc.CreateSource(pipeline, stage)

				continue
			}

			return fmt.Errorf("processor for stage %s cannot be used as source", stage.Type())

		case len(pipeline.Stages) == idx+1: //last is sink
			proc := d.processorRegistry.Get(stage)
			if proc == nil {
				return fmt.Errorf("processor for stage %s not found", stage.Type())
			}

			if sinkProc, ok := proc.(processor.Sink); ok {
				processorProvider.Sink = sinkProc.CreateSink(pipeline, stage)

				continue
			}

			return fmt.Errorf("processor for stage %s cannot be used as source", stage.Type())

		default: // others are mapper
			proc := d.processorRegistry.Get(stage)
			if proc == nil {
				return fmt.Errorf("processor for stage %s not found", stage.Type())
			}

			if mapProc, ok := proc.(processor.Map); ok {
				processorProvider.AddMapper(mapProc.CreateMap(pipeline, stage))

				continue
			}

			return fmt.Errorf("processor for stage %s cannot be used as mapper", stage.Type())
		}
	}

	wl := workload.New(processorProvider)

	errChan, err := wl.Run(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for e := range errChan {
			logger.Error(ctx, "there is an error", zap.Error(e.Err))
		}
		wg.Done()
	}()

	wl.Wait()
	wg.Wait()
	logger.Debug(ctx, "finished!!!")

	return nil
}
