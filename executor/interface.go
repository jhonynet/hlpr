package executor

import (
	"context"
	"github.com/jhonynet/hlpr/pipeline"
)

type Executor interface {
	Execute(context.Context, *pipeline.Pipeline) error
}
