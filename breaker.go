package callwrapper

import (
	"context"
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

type BreakerInterface interface {
	Execute(req func() (interface{}, error)) (interface{}, error)
	State() gobreaker.State
	Counts() gobreaker.Counts
}

type breakerStruct struct {
	breaker *gobreaker.CircuitBreaker
}

func newBreaker(cfg CallwrapperConfig) BreakerInterface {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        fmt.Sprintf("%s-%s-cb", cfg.ServiceName, cfg.FuncName),
		MaxRequests: 5,
		Timeout:     time.Duration(cfg.CBOpenToHalfOpenDuration) * time.Millisecond,
		Interval:    time.Duration(cfg.CBInterval) * time.Millisecond,
		IsSuccessful: func(err error) bool {
			if err == nil || err == context.Canceled {
				return true
			}
			if cfg.mapErrWhitelist[err] {
				return true
			}

			//only return false if err not in whitelist
			return false
		},
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			if cfg.CBOpenThreshold == 0 {
				return false
			}
			return counts.ConsecutiveFailures > uint32(cfg.CBOpenThreshold)
		},
		OnStateChange: cfg.OnCBStateChanges,
	})
	return &breakerStruct{
		breaker: cb,
	}
}

func (b *breakerStruct) Execute(req func() (interface{}, error)) (interface{}, error) {
	return b.breaker.Execute(req)
}

func (b *breakerStruct) State() gobreaker.State {
	return b.breaker.State()
}

func (b *breakerStruct) Counts() gobreaker.Counts {
	return b.breaker.Counts()
}
