package main

import "sync"

// Counter implements a simple counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments counter value
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++

}

// Value returns counter value
func (c *Counter) Value() int {
	return c.value
}

func main() {

}
