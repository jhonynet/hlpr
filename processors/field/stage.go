package field

import "github.com/jhonynet/hlpr/stages"

const (
	StageIdentifier = "field"
)

// Processor stage params.
type Stage struct {
	Name string
	Stages []*stages.Stage
}
