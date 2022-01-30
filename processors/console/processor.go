package console

import (
	"github.com/jhonynet/hlpr/pipeline"
)

const StageIdentifier = "console-output"

type Processor struct{}

func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Name() string {
	return StageIdentifier
}
