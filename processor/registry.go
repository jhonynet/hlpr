package processor

import "github.com/jhonynet/hlpr/stages"

type Registry []Processor

func (r Registry) Get(stage *stages.Stage) Processor {
	for _, p := range r {
		if p.Accepts(stage) {
			return p
		}
	}

	return nil
}
