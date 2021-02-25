package common

import (
	"sync"
)

type LightCollection struct {
	Elements map[Light]struct{}
	sync.RWMutex
}

func NewLightCollection() *LightCollection {
	return &LightCollection{Elements: make(map[Light]struct{})}
}

func (c *LightCollection) Put(el Drawable) {
	c.Lock()
	c.Elements[el] = struct{}{}
	c.Unlock()
}

func (c *LightCollection) Remove(el Drawable) {
	c.Lock()
	delete(c.Elements, el)
	c.Unlock()
}
