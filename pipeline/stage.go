package pipeline

import (
	"github.com/mitchellh/mapstructure"
)

// Stage struct represents a step in a workflow.
// stage has 3 representations
// namespace like, file::read::data.json
type Stage []interface{}

// Type return the stage type.
func (s Stage) Type() string {
	return s["type"].(string)
}

// Bind current stage to a struct.
func (s Stage) Bind(output interface{}) error {
	return mapstructure.Decode(s, output)
}
