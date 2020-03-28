package graphics

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func Black() *Color {
	return &Color{0, 0, 0, 1}
}

func White() *Color {
	return &Color{1, 1, 1, 1}
}

func Green() *Color {
	return &Color{0, 1, 0, 1}
}
