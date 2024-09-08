package main

import "fmt"

type WiredAdapter struct {}

func (w *WiredAdapter) Connect() {
	fmt.Println("Connecting by wired")
}