package modules

import (
	"fmt"
	"github.com/Gregmus2/simple-engine/common"
	"github.com/Gregmus2/simple-engine/graphics"
)

type fps struct {
	font   *graphics.Font
	frames int64
	x, y   float32
}

func (f *fps) update(dt int64) {
	if dt == 0 {
		return
	}

	f.frames = 1000 / dt
}

func (f *fps) Draw(_, _, _ float32) error {
	graphics.Text(
		f.x, f.y, 1,
		fmt.Sprintf("%d", f.frames),
		f.font, graphics.White(),
	)

	return nil
}

func NewFPS(scene common.Scene) common.UpdateActionOut {
	if !common.Config.Debug.FPS {
		return common.UpdateActionOut{Action: func(_ int64) {}}
	}

	f := &fps{
		font: graphics.GetFont("arial", 14),
		x:    float32(common.Config.Window.W - 15),
		y:    15,
	}
	scene.Drawable().Put(f)

	return common.UpdateActionOut{
		Action: f.update,
	}
}
