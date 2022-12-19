package graphics

import "github.com/Gregmus2/simple-engine/common"

func PosToUnitsX(x float32) float32 {
	return x/float32(common.Config.Window.W/2) - 1
}

func PosToUnitsY(y float32) float32 {
	return y/float32(common.Config.Window.H/2) - 1
}

func PosToUnitsW(w float32) float32 {
	return w / float32(common.Config.Window.W/2)
}

func PosToUnitsH(h float32) float32 {
	return h / float32(common.Config.Window.H/2)
}
