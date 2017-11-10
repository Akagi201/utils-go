// Package broadcast
package broadcast

import (
	"sync"
)

// Broadcast allows to send a msg to all listeners
type Broadcast struct {
	lock sync.RWMutex
	ch   chan interface{}
}

// NewBroadcast creates a new broadcast
func NewBroadcast() *Broadcast {
	return &Broadcast{
		lock: sync.RWMutex{},
		ch:   make(chan interface{}),
	}
}

// Receive a channel on which the next (close) signal will be sent
func (b *Broadcast) Receive() <-chan interface{} {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.ch
}

// Send a msg to all listeners
func (b *Broadcast) Send(msg interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.ch <- msg
}
