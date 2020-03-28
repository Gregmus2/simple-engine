package physics

import (
	"github.com/ByteArena/box2d"
	"window/common"
)

func NewWorld(cfg *common.Config) *box2d.B2World {
	gravity := box2d.MakeB2Vec2(cfg.Physics.Gravity.X, cfg.Physics.Gravity.Y)
	world := box2d.MakeB2World(gravity)

	return &world
}
