package graphics

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
)

func NewWindow(cfg *common.Config) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize glfw")
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	// for multisampling
	glfw.WindowHint(glfw.Samples, 4)
	window, err := glfw.CreateWindow(cfg.Window.W, cfg.Window.H, cfg.Window.Title, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create window")
	}
	window.MakeContextCurrent()

	window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	if glfw.RawMouseMotionSupported() {
		window.SetInputMode(glfw.RawMouseMotion, glfw.True)
	}

	return window, nil
}
