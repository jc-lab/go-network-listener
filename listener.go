package go_network_listener

import "io"

type EventType int

const (
	EventTypeAdd EventType = iota
	EventTypeRemove
	EventTypeModify
)

type AfSpec int

const (
	AfUnspecific AfSpec = 0
	AfInet4      AfSpec = 1
	AfInet6      AfSpec = 2
)

type Options struct {
	AfSpec AfSpec
}

type Event struct {
	Type EventType
}

type Listener interface {
	io.Closer
	Chan() chan *Event
}

func NewListener(options *Options) (Listener, error) {
	if options == nil {
		options = &Options{}
	}
	return newListener(options)
}
