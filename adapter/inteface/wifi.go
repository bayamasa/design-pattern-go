package main

import "fmt"

type WifiAdapter struct {}

func (w *WifiAdapter) Connect() {
	fmt.Println("Connecting by wifi")
}
