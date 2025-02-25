//go:build !windows
// +build !windows

package go_network_listener

import "errors"

func newListener(options *Options) (Listener, error) {
	return nil, errors.New("not supported platform")
}
