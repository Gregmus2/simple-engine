package graphics

import "github.com/Gregmus2/simple-engine/common"

func PercentToPosX(x float64) float64 {
	return x / 100 * float64(common.Config.Window.W)
}

func PercentToPosY(y float64) float64 {
	return y / 100 * float64(common.Config.Window.H)
}

func PercentToPosRadius(r float32) float32 {
	return r / 100 * float32(common.Config.Window.H)
}
