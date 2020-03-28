package common

type Drawable interface {
	Draw(scale float32) error
}

type Scene interface {
	Init()
	Update()
	Drawable() []Drawable
}
