package singleflight

import (
	"sync"
)

// Call is an in-flight or completed Do call
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex // protects m
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()         // wait if the request is in progress.
		return c.val, c.err // request over, return results
	}

	c := new(call)
	c.wg.Add(1)  // locker plus 1
	g.m[key] = c // there is a request to handle key
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done() // locker decrease 1

	g.mu.Lock()
	delete(g.m, key) // renew g.m
	g.mu.Unlock()

	return c.val, c.err
}
