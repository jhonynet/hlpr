package unwrap

import (
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
)

const (
	StageIdentifier = "unwrap"
)

var _ processor.Map = (*Processor)(nil)

type Processor struct {
	pipeline pipeline.Pipeline
	stage    pipeline.Stage
}

func (r *Processor) Name() string {
	return StageIdentifier
}

func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}
