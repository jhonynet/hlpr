package workload

import (
	"context"
	"fmt"
	"github.com/jhonynet/hlpr/utils/logger"
	"sync"

	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/unit"
	"github.com/jhonynet/hlpr/utils"
)

type Workload struct {
	processorProvider *processor.Provider
	dataChannel       <-chan *unit.Data
	errorChannels     []<-chan unit.Error
	waitGroup         sync.WaitGroup
}

func New(provider *processor.Provider) *Workload {
	return &Workload{
		processorProvider: provider,
	}
}

func (w *Workload) Run(ctx context.Context) (<-chan unit.Error, error) {
	if err := w.runSource(ctx); err != nil {
		return nil, err
	}

	if err := w.runMappers(ctx); err != nil {
		return nil, err
	}

	if err := w.runSink(ctx); err != nil {
		return nil, err
	}

	return utils.MergeErrors(w.errorChannels...), nil
}

func (w *Workload) Wait() {
	w.waitGroup.Wait()
}

func (w *Workload) runSource(ctx context.Context) error {
	var (
		errChan <-chan unit.Error
		err     error
	)
	if w.dataChannel, errChan, err = w.processorProvider.Source.RunSource(ctx, &w.waitGroup); err != nil {
		return fmt.Errorf(
			"failed to setup %s as source processor: %w",
			w.processorProvider.Source.Identifier(), err,
		)
	}
	logger.Debug(ctx, fmt.Sprintf(
		"source processor %s running",
		w.processorProvider.Source.Identifier(),
	))

	w.errorChannels = append(w.errorChannels, errChan)

	return nil
}

func (w *Workload) runMappers(ctx context.Context) error {
	for _, mapper := range w.processorProvider.Mappers {
		var (
			errChan <-chan unit.Error
			err     error
		)
		if w.dataChannel, errChan, err = mapper.RunMap(ctx, w.dataChannel, &w.waitGroup); err != nil {
			return fmt.Errorf(
				"failed to setup %s as mapper processor: %w",
				mapper.Identifier(), err,
			)
		}
		logger.Debug(ctx, fmt.Sprintf(
			"mapper processor %s running",
			mapper.Identifier(),
		))

		w.errorChannels = append(w.errorChannels, errChan)
	}

	return nil
}

func (w *Workload) runSink(ctx context.Context) error {
	errChan, err := w.processorProvider.Sink.RunSink(ctx, w.dataChannel, &w.waitGroup)
	if err != nil {
		return fmt.Errorf(
			"failed to setup %s as sink processor: %w",
			w.processorProvider.Sink.Identifier(), err,
		)
	}
	logger.Debug(ctx, fmt.Sprintf(
		"sink processor %s running",
		w.processorProvider.Sink.Identifier(),
	))

	w.errorChannels = append(w.errorChannels, errChan)

	return nil
}
