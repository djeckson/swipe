//+build !swipe

// Code generated by Swipe v1.20.0. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"context"
	"github.com/go-kit/kit/metrics"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"time"
)

type instrumentingMiddlewareServiceInterface struct {
	next           service.Interface
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

func (s *instrumentingMiddlewareServiceInterface) Delete(ctx context.Context, id uint) (a string, b string, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Delete").Add(1)
		s.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Delete(ctx, id)
}

func (s *instrumentingMiddlewareServiceInterface) Get(ctx context.Context, id int, name string, fname string, price float32, n int) (data user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Get").Add(1)
		s.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Get(ctx, id, name, fname, price, n)
}

func (s *instrumentingMiddlewareServiceInterface) GetAll(ctx context.Context) (_ []*user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetAll").Add(1)
		s.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.GetAll(ctx)
}

func (s *instrumentingMiddlewareServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (states map[string]map[int][]string, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TestMethod").Add(1)
		s.requestLatency.With("method", "TestMethod").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.TestMethod(data, ss)
}

func (s *instrumentingMiddlewareServiceInterface) Create(ctx context.Context, name string, data []byte) (_ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Create").Add(1)
		s.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Create(ctx, name, data)
}

func NewInstrumentingMiddlewareServiceInterface(s service.Interface, requestCount metrics.Counter, requestLatency metrics.Histogram) service.Interface {
	return &instrumentingMiddlewareServiceInterface{next: s, requestCount: requestCount, requestLatency: requestLatency}
}
