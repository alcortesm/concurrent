package concurrent

import "sync"

// Flag is a concurrent-safe Boolean value. It starts unset and once it
// has been set, it cannot be unset again.
type Flag struct {
	ch   chan struct{}
	once sync.Once
}

// NewFlag returns a new (unset) flag ready to use.
func NewFlag() *Flag {
	f := Flag{
		ch: make(chan struct{}),
	}

	return &f
}

// Set sets the flag permanently. Setting an already set flag is safe
// but doesn't have any effect on the flag.
func (f *Flag) Set() {
	f.once.Do(func() {
		close(f.ch)
	})
}

// IsSet returns true if the flag is set or false otherwise.
func (f *Flag) IsSet() bool {
	select {
	case <-f.ch:
		return true
	default:
		return false
	}
}

// Done returns a channel that's closed  when the flag is Set.
// Successive calls to Done return the same value. The close of the Done
// channel may happen asynchronously, after the Set method returns.
//
// Done is provided to use in select statements, similarly to how
// you would use the Done method of the stdlib context.Context type.
func (f *Flag) Done() <-chan struct{} {
	return f.ch
}
