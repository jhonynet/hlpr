package runner

import "context"

type Runner interface {
	Run(context.Context) error
}

type runner struct {

}

func (r *runner) Run(ctx context.Context) error {
	return nil
}

func NewRunner() Runner {
	return &runner{}
}
