package stages

import (
	"github.com/mitchellh/mapstructure"
)

// This struct represents a step in a workflow.
type Stage map[string]interface{}

// Return the stage type.
func (s Stage) Type() string {
	return s["type"].(string)
}

// Bind current stage to an struct.
func (s Stage) Bind(output interface{}) error {
	return mapstructure.Decode(s, output)
}
