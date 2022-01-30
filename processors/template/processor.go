package template

import (
	"github.com/jhonynet/hlpr/pipeline"
)

type Processor struct {
	stage    pipeline.Stage
	pipeline pipeline.Pipeline
}

func (r *Processor) Name() string {
	return StageIdentifier
}

func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}
