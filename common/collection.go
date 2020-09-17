package common

import "sync"

type DrawableCollection struct {
	Elements map[Drawable]struct{}
	sync.RWMutex
}

func NewDrawableCollection() *DrawableCollection {
	return &DrawableCollection{Elements: make(map[Drawable]struct{})}
}

func (c *DrawableCollection) Put(el Drawable) {
	c.Lock()
	c.Elements[el] = struct{}{}
	c.Unlock()
}

func (c *DrawableCollection) Remove(el Drawable) {
	c.Lock()
	delete(c.Elements, el)
	c.Unlock()
}
