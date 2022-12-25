package modules

import (
	"github.com/Gregmus2/simple-engine/common"
	"github.com/sirupsen/logrus"
)

func NewFPS(logger *logrus.Entry) common.UpdateActionOut {
	return common.UpdateActionOut{
		Action: func(dt int64) {
			if !common.Config.Debug.FPS {
				return
			}

			logger.Debug(1 / dt)
		},
	}
}
