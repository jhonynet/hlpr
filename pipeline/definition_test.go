package pipeline

import (
	"errors"
	"testing"

	"github.com/jhonynet/hlpr/stages"
	"github.com/stretchr/testify/assert"
)

func TestFromBytes(t *testing.T) {
	tests := []struct {
		name   string
		bytes  []byte
		format string
		want   Definition
		err    error
	}{
		{
			name:   "encode json",
			bytes:  getJsonDefinition(),
			format: ".json",
			want: Definition{
				Vars: nil,
				Stages: []stages.Stage{
					{
						"type": "raw-input",
						"data": []interface{}{"jhony", "carlos"},
					},
				},
			},
		},
		{
			name:   "encode yaml",
			bytes:  getYamlDefinition(),
			format: ".yaml",
			want: Definition{
				Vars: map[string]interface{}{
					"configA": "is valid",
				},
				Stages: []stages.Stage{
					{
						"type": "raw-input",
						"data": []interface{}{"jhony", "carlos"},
					},
				},
			},
		},
		{
			name:   "invalid format",
			format: ".xml",
			want:   Definition{},
			err:    errors.New("cannot parse workflow filename => format .xml is not supported, only json or yaml"),
		},
		{
			name:   "error unmarshaling",
			bytes:  getYamlDefinition(),
			format: ".json",
			want:   Definition{},
			err:    errors.New("cannot parse definition file => invalid character 'v' looking for beginning of value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromBytes(tt.bytes, tt.format)
			assert.Equal(t, tt.want, got)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func getYamlDefinition() []byte {
	return []byte(`
vars:
  configA: is valid
stages:
  - type: raw-input
    data:
      - jhony
      - carlos
`)
}

func getJsonDefinition() []byte {
	return []byte(`{
  "stages": [
    {
      "type": "raw-input",
      "data": [
        "jhony",
        "carlos"
      ]
    }
  ]
}`)
}
