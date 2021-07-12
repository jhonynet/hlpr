package processors

import (
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/processors/console"
	"github.com/jhonynet/hlpr/processors/file"
	"github.com/jhonynet/hlpr/processors/http"
	"github.com/jhonynet/hlpr/processors/jq"
	"github.com/jhonynet/hlpr/processors/json"
	"github.com/jhonynet/hlpr/processors/raw"
	"github.com/jhonynet/hlpr/processors/template"
	"github.com/jhonynet/hlpr/processors/unwrap"
)

var DefaultRegistry = processor.Registry{
	new(raw.Processor),
	new(template.Processor),
	new(file.Processor),
	new(http.Processor),
	new(jq.Processor),
	new(json.Processor),
	new(unwrap.Processor),
	new(console.Processor),
}
