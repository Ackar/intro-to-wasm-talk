package main

import (
	"syscall/js"
	"time"
)

func main() {
	counter := js.Global().Get("timestamp")
	for {
		counter.Set("innerHTML", time.Now().Unix())
		time.Sleep(1 * time.Second)
	}
}
