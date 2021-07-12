package pipeline

type Pipeline struct {
	Definition Definition
}

func NewPipeline(definition Definition) *Pipeline {
	return &Pipeline{
		Definition: definition,
	}
}
