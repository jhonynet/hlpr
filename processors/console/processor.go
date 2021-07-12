package console

import (
	"github.com/jhonynet/hlpr/stages"
)

const StageIdentifier = "console-output"

type Processor struct{}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}
