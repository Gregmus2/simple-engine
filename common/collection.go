package common

type DrawableCollection struct {
	Elements map[Drawable]struct{}
}

func NewDrawableCollection() *DrawableCollection {
	return &DrawableCollection{Elements: make(map[Drawable]struct{})}
}

func (c *DrawableCollection) Put(el Drawable) {
	c.Elements[el] = struct{}{}
}

func (c *DrawableCollection) Remove(el Drawable) {
	delete(c.Elements, el)
}
