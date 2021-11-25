package processor

import (
	"github.com/jhonynet/hlpr/pipeline"
)

type Registry []Processor

func (r Registry) Get(stage pipeline.Stage) Processor {
	for _, p := range r {
		if p.Accepts(stage) {
			return p
		}
	}

	return nil
}
