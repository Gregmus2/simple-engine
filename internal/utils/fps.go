package utils

import (
	"fmt"
	"github.com/Gregmus2/simple-engine/internal/dispatchers"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type FPS struct {
	window *glfw.Window

	frames  int
	seconds float64
}

func NewFPS(w *glfw.Window, u *dispatchers.Update) *FPS {
	f := &FPS{window: w}

	u.Subscribe(f.Update)

	return f
}

func (f *FPS) Update(delta float64) {
	f.frames++
	f.seconds += delta
	if f.seconds >= 1.0 {
		f.window.SetTitle(fmt.Sprintf("[%d FPS]", f.frames))
		f.seconds = 0
		f.frames = 0
	}
}
