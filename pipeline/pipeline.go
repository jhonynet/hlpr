package pipeline

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	Metadata Metadata `yaml:"metadata"`
	Stages   []Stage  `yaml:"stages"`
}

func FromYaml(bytes []byte) (Pipeline, error) {
	var pipeline Pipeline
	if err := yaml.Unmarshal(bytes, &pipeline); err != nil {
		return pipeline, fmt.Errorf("cannot unmarshal pipeline from yaml bytes => %w", err)
	}

	return pipeline, nil
}
