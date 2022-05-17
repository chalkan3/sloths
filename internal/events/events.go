package events

type IEvent interface {
	String()
}

type Event struct{}

func (e Event) String() {}
