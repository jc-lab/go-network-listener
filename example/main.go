package main

import (
	go_network_listener "github.com/jc-lab/go-network-listener"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	listener, err := go_network_listener.NewListener(nil)
	if err != nil {
		log.Fatalf("new failed: %+v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		log.Printf("CLOSE")
		listener.Close()
	}()

	for {
		event, ok := <-listener.Chan()
		if !ok {
			break
		}
		log.Printf("EVENT: %+v", event)
	}
}
