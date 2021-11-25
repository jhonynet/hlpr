package jq

import (
	"github.com/jhonynet/hlpr/pipeline"
)

// Raw Input Processor allows to inject custom data into the workflow.
type Processor struct {
	pipeline pipeline.Pipeline
	stage    pipeline.Stage
}

func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}
