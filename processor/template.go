package processor

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/jhonynet/hlpr/unit"
	"io"
	"text/template"
)

type Template struct {
	underlying *template.Template
}

func NewTemplate(tpl string, vars map[string]interface{}) (*Template, error) {
	parsed, err := template.
		New("").
		Funcs(sprig.TxtFuncMap()).
		Funcs(getVarsFunc(vars)).
		Parse(tpl)

	if err != nil {
		return nil, err
	}

	return &Template{
		underlying: parsed,
	}, nil
}

func (t *Template) WithVars(vars map[string]interface{}) *Template {
	t.underlying = t.underlying.Funcs(template.FuncMap{
		"vars": func(key string) string {
			if value, ok := vars[key]; ok {
				return value.(string)
			}

			return ""
		},
	})

	return t
}

func (t *Template) Render(wr io.Writer, unit *unit.Data) error {
	return t.underlying.Execute(wr, unit.Value)
}

// getVarsFunc return internal func map.
func getVarsFunc(vars map[string]interface{}) template.FuncMap {
	if vars == nil {
		return nil
	}

	return template.FuncMap{
		"vars": func(key string) string {
			if value, ok := vars[key]; ok {
				return value.(string)
			}

			return ""
		},
	}
}
