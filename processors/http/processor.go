package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processor"
	"github.com/jhonynet/hlpr/stages"
	"github.com/jhonynet/hlpr/unit"
	"github.com/jhonynet/hlpr/utils"
	"github.com/jhonynet/hlpr/utils/logger"
	"io/ioutil"
	"net/http"
	"sync"
)

type Processor struct {
	pipeline *pipeline.Pipeline
	stage    *stages.Stage
}

func (r *Processor) Accepts(s *stages.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Identifier() string {
	return StageIdentifier
}

// todo: follow context cancellation signal.
func (r *Processor) runWorkers(ctx context.Context, stage Stage, urlTpl, bodyTpl *processor.Template, input <-chan *unit.Data, output chan *unit.Data, errChan chan unit.Error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	defer close(errChan)

	var wp = r.getWorkerPool(stage)
	for data := range input {
		wp.Work(func(data *unit.Data) func() {
			return func() {
				res, err := r.processUnit(ctx, stage, urlTpl, bodyTpl, data)
				if err != nil {
					errChan <- unit.Error{Err: err}
					return
				}

				output <- res
			}
		}(data))
	}
	wp.Wait()
	wp.Close()
}

// parse input data and make an http request.
func (r *Processor) processUnit(ctx context.Context, stage Stage, urlTpl, bodyTpl *processor.Template, data *unit.Data) (*unit.Data, error) {
	url, err := r.renderTemplate(urlTpl, data)
	if err != nil {
		return nil, err
	}

	body, err := r.renderTemplate(bodyTpl, data)
	if err != nil {
		return nil, err
	}

	headers, err := r.renderHeaders(stage, data)
	if err != nil {
		return nil, err
	}

	method := http.MethodGet
	if stage.Method != "" {
		method = stage.Method
	}

	logger.Debug(ctx, fmt.Sprintf("Requesting a %s to URL %s with body %s", method, url.String(), body.String()))
	req, err := http.NewRequest(method, url.String(), &body)
	if err != nil {
		return nil, err
	}

	req.Header = headers

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	out, err := r.parseBody(stage, resp)
	if err != nil {
		return nil, err
	}

	return data.SetValue(out), nil
}

// todo: use content type from stage.
func (r *Processor) parseBody(_ Stage, res *http.Response) (interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Processor) renderHeaders(stage Stage, data *unit.Data) (http.Header, error) {
	out := http.Header{}
	for key, template := range stage.Headers {
		// todo: if this is not a template just set it.
		var value bytes.Buffer
		headerTpl, err := processor.NewTemplate(template, r.pipeline.Definition.Vars)
		if err != nil {
			return nil, fmt.Errorf("cannot create template for header %s:%s => %w", key, template, err)
		}
		if err := headerTpl.Render(&value, data); err != nil {
			return nil, fmt.Errorf("cannot render template for header %s:%s => %w", key, template, err)
		}
		out.Set(key, value.String())
	}

	return out, nil
}

func (r *Processor) renderTemplate(tpl *processor.Template, data *unit.Data) (out bytes.Buffer, err error) {
	err = tpl.Render(&out, data)

	return
}

func (r *Processor) getWorkerPool(stage Stage) *utils.WorkerPool {
	threads := 1
	if stage.Threads != 0 {
		threads = stage.Threads
	}

	return utils.NewWorkerPool(threads)
}
