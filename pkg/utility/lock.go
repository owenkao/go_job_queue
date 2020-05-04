package service

import (
	"errors"
	"time"
)

// Locker ...
type Locker struct {
	timeout time.Duration
	queue   chan struct{}
}

// NewLocker create new locker instance.
func NewLocker(timeout time.Duration) *Locker {
	return &Locker{
		timeout,
		make(chan struct{}, 1),
	}
}

// Lock lock with timeout
func (l *Locker) Lock() error {
	select {
	case l.queue <- struct{}{}:
		return nil
	case <-time.After(time.Duration(l.timeout) * time.Second):
		return errors.New("timeout")
	}
}

// Unlock release lock
func (l *Locker) Unlock() {
	<-l.queue
}
