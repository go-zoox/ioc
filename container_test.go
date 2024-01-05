package ioc

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Error("New() should not return nil")
	}

	c.Register("test", 1)
	if c.Has("test") != true {
		t.Error("Has() should return true")
	}

	service, ok := c.Get("test")
	if !ok {
		t.Error("Get() should return true")
	} else if service != 1 {
		t.Error("Get() should return 1")
	}

	if c.MustGet("test") != 1 {
		t.Error("MustGet() should return 1")
	}
}

func TestInvoke(t *testing.T) {
	x := 0
	c := New()
	c.Register("test", func(a int) int { return a + 1 })
	c.Invoke("test", func(fn func(a int) int, ok bool) {
		x = fn(1)
	})

	if x != 2 {
		t.Error("Invoke() should return 2")
	}
}

func TestMustInvolke(t *testing.T) {
	x := 0
	c := New()
	c.Register("test", func(a int) int { return a + 1 })
	c.MustInvoke("test", func(fn func(a int) int) {
		x = fn(1)
	})

	if x != 2 {
		t.Error("Invoke() should return 2")
	}
}
