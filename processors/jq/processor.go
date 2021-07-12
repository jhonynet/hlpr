package jq

import (
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/stages"
)

// Raw Input Processor allows to inject custom data into the workflow.
type Processor struct {
	pipeline *pipeline.Pipeline
	stage    *stages.Stage
}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}
