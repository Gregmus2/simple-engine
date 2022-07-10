package common

import "go.uber.org/dig"

type UpdateActionOut struct {
	dig.Out
	Action func() `group:"update_actions"`
}

type UpdateActionsIn struct {
	dig.In
	Actions []func() `group:"update_actions"`
}

type Pos struct {
	X, Y float32
}
