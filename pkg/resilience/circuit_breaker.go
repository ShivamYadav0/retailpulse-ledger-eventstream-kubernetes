package resilience

import (
    "errors"
    "sync"
    "time"
)

const (
    stateClosed   = "closed"
    stateOpen     = "open"
    stateHalfOpen = "half_open"
)

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitBreaker struct {
    mu               sync.Mutex
    failures         int
    state            string
    lastFailure      time.Time
    failureThreshold int
    resetTimeout     time.Duration
}

func NewCircuitBreaker(failureThreshold int, resetTimeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:            stateClosed,
        failureThreshold: failureThreshold,
        resetTimeout:     resetTimeout,
    }
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    cb.mu.Lock()
    if cb.state == stateOpen {
        if time.Since(cb.lastFailure) < cb.resetTimeout {
            cb.mu.Unlock()
            return ErrCircuitOpen
        }
        cb.state = stateHalfOpen
    }
    cb.mu.Unlock()

    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.failureThreshold {
            cb.state = stateOpen
        }
        return err
    }

    cb.failures = 0
    cb.state = stateClosed
    return nil
}

func (cb *CircuitBreaker) State() string {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    return cb.state
}
