package unwrap

import (
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
)

const (
	StageIdentifier = "unwrap"
)

var _ processor.Map = (*Processor)(nil)

type Processor struct {
	pipeline *pipeline.Pipeline
	stage    *stages.Stage
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}
