package promise

import (
	"context"
	"sync"
)

type Map struct {
	fn func(context.Context, interface{}) (interface{}, error)
	m  *sync.Map
}

func NewMap(fn func(context.Context, interface{}) (interface{}, error)) *Map {
	return &Map{fn, new(sync.Map)}
}

func (m *Map) Get(ctx context.Context, key, data interface{}) (interface{}, error) {
	p, _ := m.m.LoadOrStore(key, New(func(ctx context.Context) (interface{}, error) {
		res, err := m.fn(ctx, data)
		if err != nil {
			m.m.Delete(key)
		}
		return res, err
	}))
	return p.(*Promise).Await(ctx)
}
