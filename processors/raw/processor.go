package raw

import (
	"github.com/jhonynet/hlpr/pipeline"
)

type Processor struct {
	stage pipeline.Stage
}

func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Name() string {
	return StageIdentifier
}
