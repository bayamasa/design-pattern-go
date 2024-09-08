package main

import "fmt"

type WifiAdapter struct {
	wiredAdapter *WiredAdapter
}

func NewWifiAdapter() *WifiAdapter {
	wiredAdapter := NewWiredAdapter("localhost")
	return &WifiAdapter{wiredAdapter: wiredAdapter}
}

func (w *WifiAdapter) Connect() {
	fmt.Println("Connecting by wifi to", w.wiredAdapter.GetHost())
}
