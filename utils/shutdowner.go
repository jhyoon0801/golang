package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type ShutDowner struct {
	sigs     chan os.Signal
	callback func(...interface{}) error
}

func (shutDowner *ShutDowner) Init() {
	shutDowner.sigs = make(chan os.Signal, 1)
	signal.Notify(shutDowner.sigs, syscall.SIGINT, syscall.SIGTERM)
}

func (shutDowner *ShutDowner) SetCallback(callback func(...interface{}) error) {
	shutDowner.callback = callback
}

func (shutDowner *ShutDowner) Shutdown() {
	shutDowner.sigs <- syscall.SIGINT
}

func (shutDowner *ShutDowner) WaitSignal(param ...interface{}) error {
	<-shutDowner.sigs
	return shutDowner.callback(param)
}
