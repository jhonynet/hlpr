package raw

import (
	"github.com/jhonynet/hlpr/stages"
)

type Processor struct {
	stage *stages.Stage
}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}
