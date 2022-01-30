package processor

import (
	"context"
)

type Collection []Processor

// Processor is a generic way to process data stream.
type Processor interface {
	Process(context.Context)
}
