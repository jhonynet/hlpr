package pipeline

import (
	"encoding/json"
	"fmt"

	"github.com/jhonynet/hlpr/stages"
	"gopkg.in/yaml.v2"
)

type Definition struct {
	Vars   map[string]interface{}
	Stages []*stages.Stage
}

func FromBytes(bytes []byte, format string) (Definition, error) {
	var (
		def           Definition
		unmarshalFunc func([]byte, interface{}) error
	)

	switch format {
	case ".json":
		unmarshalFunc = json.Unmarshal
		break
	case ".yml", ".yaml":
		unmarshalFunc = yaml.Unmarshal
		break
	default:
		return def, fmt.Errorf(
			"cannot parse workflow filename => format %s is not supported, only json or yaml", format,
		)
	}

	if err := unmarshalFunc(bytes, &def); err != nil {
		return def, fmt.Errorf("cannot parse definition file => %w", err)
	}

	return def, nil
}
