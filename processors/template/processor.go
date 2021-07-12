package template

import (
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/stages"
)

type Processor struct {
	stage    *stages.Stage
	pipeline *pipeline.Pipeline
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}
