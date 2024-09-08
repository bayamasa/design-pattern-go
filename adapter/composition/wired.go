package main

import "fmt"

type WiredAdapter struct {
	host string
}

func NewWiredAdapter(host string) *WiredAdapter {
	return &WiredAdapter{host: host}
}

func (w *WiredAdapter) Connect() {
	fmt.Println("Connecting by wired to", w.host)
}

func (w *WiredAdapter) GetHost() string {
	return w.host
}