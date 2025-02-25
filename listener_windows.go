package go_network_listener

import (
	"github.com/jc-lab/go-network-listener/internal/iphlpapi"
	"github.com/jc-lab/go-network-listener/internal/retainer"
	"golang.org/x/sys/windows"
	"sync"
	"unsafe"
)

type windowsListener struct {
	mutex                  sync.Mutex
	closed                 bool
	mibChangeNotify2Handle windows.Handle

	events chan *Event
}

var notifyIpInterfaceChangeCallback = windows.NewCallback(func(callerContext uintptr, row *iphlpapi.MibIPInterfaceRow, notificationType iphlpapi.MibNotificationType) uintptr {
	listener, ok := retainer.Get(callerContext).(*windowsListener)
	if !ok {
		return 0
	}

	event := &Event{}
	switch notificationType {
	case iphlpapi.MibParameterNotification:
		event.Type = EventTypeModify
	case iphlpapi.MibAddInstance:
		event.Type = EventTypeAdd
	case iphlpapi.MibDeleteInstance:
		event.Type = EventTypeRemove
	default:
		return 0
	}

	listener.mutex.Lock()
	defer listener.mutex.Unlock()
	if listener.closed {
		return 0
	}

	select {
	case listener.events <- event:
	default:
		// Channel is full, skip event
	}

	return 0
})

func newListener(options *Options) (Listener, error) {
	l := &windowsListener{
		events: make(chan *Event, 10),
	}
	ptr := retainer.Retain(uintptr(unsafe.Pointer(l)), l)

	err := iphlpapi.NotifyIpInterfaceChange(
		options.AfSpec.toWindows(), // Both IPv4 and IPv6
		notifyIpInterfaceChangeCallback,
		ptr,
		false,
		&l.mibChangeNotify2Handle,
	)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *windowsListener) Chan() chan *Event {
	return l.events
}

func (l *windowsListener) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.closed = true
	close(l.events)

	if l.mibChangeNotify2Handle != 0 {
		_ = windows.CancelMibChangeNotify2(l.mibChangeNotify2Handle)
		l.mibChangeNotify2Handle = 0
	}

	retainer.Release(uintptr(unsafe.Pointer(l)))

	return nil
}

func (v AfSpec) toWindows() uint16 {
	switch v {
	case AfInet4:
		return windows.AF_INET
	case AfInet6:
		return windows.AF_INET6
	default:
		return windows.AF_UNSPEC
	}
}
