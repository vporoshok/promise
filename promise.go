package promise

import (
	"context"
	"sync"
)

type Promise struct {
	fn     func(context.Context) (interface{}, error)
	once   *sync.Once
	result interface{}
	err    error
}

func New(fn func(context.Context) (interface{}, error)) *Promise {
	return &Promise{fn, new(sync.Once), nil, nil}
}

func (p *Promise) Await(ctx context.Context) (interface{}, error) {
	p.once.Do(func() {
		p.result, p.err = p.fn(ctx)
	})
	return p.result, p.err
}
