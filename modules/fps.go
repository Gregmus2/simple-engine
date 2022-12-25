package modules

import (
	"github.com/Gregmus2/simple-engine/common"
)

func NewFPS() common.UpdateActionOut {
	return common.UpdateActionOut{
		Action: func(dt int64) {
			if !common.Config.Debug.FPS {
				return
			}

			println(1 / dt)
		},
	}
}
