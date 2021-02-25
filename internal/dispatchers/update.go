package dispatchers

type Update struct {
	subscribers []func(delta float64)
}

func NewUpdate() *Update {
	return &Update{subscribers: make([]func(delta float64), 0)}
}

func (u *Update) Subscribe(subscriber func(delta float64)) {
	u.subscribers = append(u.subscribers, subscriber)
}

func (u *Update) Dispatch(delta float64) {
	for _, subscriber := range u.subscribers {
		subscriber(delta)
	}
}
