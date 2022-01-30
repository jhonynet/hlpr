package file

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/unit"
	"github.com/jhonynet/hlpr/utils/logger"
	"go.uber.org/zap"
)

type Processor struct {
	pipeline pipeline.Pipeline
	stage    pipeline.Stage
}

// Return if this processor accepts the stage.
func (r *Processor) Accepts(s pipeline.Stage) bool {
	return s.Type() == StageIdentifier
}

func (r *Processor) Name() string {
	return StageIdentifier
}

// read from file and send each line to the data steam.
// todo: cancel if context is cancelled.
func (r *Processor) read(ctx context.Context, stage Stage, f *os.File, input <-chan *unit.Data, output chan *unit.Data, errChan chan unit.Error, wg *sync.WaitGroup) {
	defer close(output)
	defer close(errChan)
	defer wg.Done()

	if input != nil {
		for data := range input {
			output <- data
		}
	}

	if stage.Split {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			output <- &unit.Data{Value: scanner.Text()}
		}
	} else {
		b, err := ioutil.ReadAll(f)
		if err != nil {
			errChan <- unit.Error{Err: err}

			return
		}
		output <- &unit.Data{Value: b}
	}

	if err := f.Close(); err != nil {
		logger.Error(ctx, "cannot close file", zap.Error(err))
	}
}

func (r Processor) write(ctx context.Context, stage Stage, f *os.File, input <-chan *unit.Data, output chan *unit.Data, errChan chan unit.Error, wg *sync.WaitGroup) {
	defer close(errChan)
	defer wg.Done()
	if output != nil {
		defer close(output)
	}

	var written int

	for data := range input {
		var (
			n   int
			err error
		)

		switch d := data.Value.(type) {
		case string:
			// todo: check why \n is writted as literal.
			n, err = f.WriteString(strings.ReplaceAll(d, `\n`, "\n"))
			break

		case []byte:
			n, err = f.Write(d)
			break

		default:
			err = fmt.Errorf("cannot write data to file: %T cannot be written into a file", data.Value)
		}

		written += n
		if err != nil {
			errChan <- unit.Error{Err: err}
			continue
		}

		if output != nil {
			output <- data
		}
	}

	if err := f.Close(); err != nil {
		logger.Error(ctx, "cannot close file", zap.Error(err))
	}
	if written == 0 {
		if err := os.Remove(stage.Path); err != nil {
			logger.Error(ctx, "cannot delete file", zap.String("file", stage.Path), zap.Error(err))
		}
	}
}
