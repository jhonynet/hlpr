package workload

import (
	"context"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/processors"
	"github.com/jhonynet/hlpr/processors/raw"
	"github.com/jhonynet/hlpr/processors/template"
	"github.com/jhonynet/hlpr/stages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicWorkload(t *testing.T) {
	rawInputProcessor := &raw.Processor{}
	templateProcessor := &template.Processor{}
	consoleOutputProcessor := &processors.ConsoleOutput{}

	reg := &processor.Provider{
		Source: rawInputProcessor.CreateSource(nil, &stages.Stage{
			"data": []string{"jhony", "marcos"},
		}),
		Mappers: []processor.Map{
			templateProcessor.CreateMap(nil, &stages.Stage{
				"output": "{{ . }} decorado por A",
			}),
			templateProcessor.CreateMap(nil, &stages.Stage{
				"output": "{{ . }}, decorado por B",
			}),
			templateProcessor.CreateMap(nil, &stages.Stage{
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
