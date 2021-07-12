package processor

import (
	"context"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/stages"
	"sync"

	"github.com/jhonynet/hlpr/unit"
)

// Processor is a generic way to process data stream.
type Processor interface {
	// Identifier is the unique id of each processor.
	Identifier() string

	// Accepts returns the id of the processor.
	Accepts(*stages.Stage) bool
}

// Source represents a processor that's generates data.
type Source interface {
	Processor

	// Create Source
	CreateSource(*pipeline.Pipeline, *stages.Stage) Source

	// Source will run this processor as source.
	RunSource(context.Context, *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error)
}

// Map takes an Input, mutate it and then provide it as an Output.
type Map interface {
	Processor

	// Create Map
	CreateMap(*pipeline.Pipeline, *stages.Stage) Map

	// Map will run this processor as mapper.
	RunMap(context.Context, <-chan *unit.Data, *sync.WaitGroup) (<-chan *unit.Data, <-chan unit.Error, error)
}

// Sink do something with the mutated data.
type Sink interface {
	Processor

	// Create Sink
	CreateSink(*pipeline.Pipeline, *stages.Stage) Sink

	// Sink will run this processor as sink.
	RunSink(context.Context, <-chan *unit.Data, *sync.WaitGroup) (<-chan unit.Error, error)
}
