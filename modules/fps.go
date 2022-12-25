package modules

import (
	"github.com/Gregmus2/simple-engine/common"
)

func NewFPS() common.UpdateActionOut {
	return common.UpdateActionOut{
		Action: func(dt int64) {
			if !common.Config.Debug.FPS || dt == 0 {
				return
			}

			println(1000 / dt)
		},
	}
}
