package json

import (
	"fmt"
	"reflect"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/unit"
	"github.com/json-iterator/go"
)

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

func (r *Processor) decode(input *unit.Data) (*unit.Data, error) {
	var in []byte
	switch val := input.Value.(type) {

	case string:
		in = []byte(val)
		break

	case []byte:
		in = val
		break

	default:
		return nil, fmt.Errorf("cannot json-unmarshal %s type", reflect.TypeOf(val).String())
	}

	var out interface{}
	if err := jsoniter.Unmarshal(in, &out); err != nil {
		return nil, fmt.Errorf("cannot json-unmarshal this value %w", err)
	}

	return input.SetValue(out), nil
}

func (r *Processor) encode(input *unit.Data) (*unit.Data, error) {
	out, err := jsoniter.Marshal(input.Value)
	if err != nil {
		return nil, fmt.Errorf("cannot json-unmarshal this value %w", err)
	}

	return input.SetValue(out), nil
}
