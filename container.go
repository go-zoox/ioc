package ioc

import (
	"fmt"
	"reflect"
	"sync"
)

// Container is DI Container
type Container interface {
	Register(id string, service any)
	Get(id string) (any, bool)
	MustGet(id string) any
	Invoke(id string, fn any)
	MustInvoke(id string, fn any)
	Has(id string) bool
	//
	Length() int
	ForEach(fn func(id string, service any))
}

type container struct {
	sync.RWMutex
	registry map[string]any
}

// Register service by id
func (c *container) Register(id string, service any) {
	c.Lock()
	defer c.Unlock()

	c.registry[id] = service
}

// Get gets a service by id
func (c *container) Get(id string) (any, bool) {
	c.RLock()
	defer c.RUnlock()

	service, ok := c.registry[id]
	return service, ok
}

// MustGet calls Get underneath
// will panic if serviceect not found within container
func (c *container) MustGet(id string) any {
	service, ok := c.Get(id)
	if !ok {
		panic(fmt.Sprintf("Service <%s> not found in container", id))
	}

	return service
}

// Invoke gets a service safely typed by passing it to a closure
// will panic if callback is not a function
func (c *container) Invoke(id string, fn any) {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic(fmt.Sprintf("<%s> is not a reflect.Func", reflect.TypeOf(fn)))
	}

	o, ok := c.Get(id)
	callback := reflect.ValueOf(fn)
	args := []reflect.Value{reflect.ValueOf(o), reflect.ValueOf(ok)}

	callback.Call(args)
}

// MustInvoke calls MustGet underneath
// will panic if service not found within container
func (c *container) MustInvoke(id string, fn any) {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic(fmt.Sprintf("<%s> is not a reflect.Func", reflect.TypeOf(fn)))
	}

	o := c.MustGet(id)
	callback := reflect.ValueOf(fn)
	args := []reflect.Value{reflect.ValueOf(o)}

	callback.Call(args)
}

// Has checks if a service exists within the container
func (c *container) Has(id string) bool {
	_, ok := c.Get(id)
	return ok
}

// Length returns the length of the container
func (c *container) Length() int {
	c.RLock()
	defer c.RUnlock()

	return len(c.registry)
}

// ForEach loops the services of the container
func (c *container) ForEach(fn func(id string, service any)) {
	c.RLock()
	defer c.RUnlock()

	for id, service := range c.registry {
		fn(id, service)
	}
}
