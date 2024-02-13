package observer

import "context"

type ISubscriber interface {
	Update(ctx context.Context)
}

type IObservable interface {
	AddSubscriber(ISubscriber)
	RemoveSubscriber(ISubscriber)
	Notify()
}
