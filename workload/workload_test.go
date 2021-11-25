package workload

import (
	"context"
	"testing"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/processors"
	"github.com/jhonynet/hlpr/processors/raw"
	"github.com/jhonynet/hlpr/processors/template"
	"github.com/stretchr/testify/assert"
)

func TestBasicWorkload(t *testing.T) {
	rawInputProcessor := &raw.Processor{}
	templateProcessor := &template.Processor{}
	consoleOutputProcessor := &processors.ConsoleOutput{}

	reg := &processor.Provider{
		Source: rawInputProcessor.CreateSource(nil, &pipeline.Stage{
			"data": []string{"jhony", "marcos"},
		}),
		Mappers: []processor.Map{
			templateProcessor.CreateMap(nil, &pipeline.Stage{
				"output": "{{ . }} decorado por A",
			}),
			templateProcessor.CreateMap(nil, &pipeline.Stage{
				"output": "{{ . }}, decorado por B",
			}),
			templateProcessor.CreateMap(nil, &pipeline.Stage{
				"output": "{{ . }} y Decorado por C",
			}),
		},
		Sink: consoleOutputProcessor.CreateSink(nil, nil),
	}

	ctx := context.Background()

	workload := New(reg)

	_, err := workload.Run(ctx)

	assert.Nil(t, err)

	workload.Wait()
}
