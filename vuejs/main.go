package main

import (
	"syscall/js"
)

func reverse(s string) string {
	var res string
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func main() {
	vue := js.Global().Get("Vue")
	m := map[string]interface{}{
		"el": "#app",
		"data": map[string]interface{}{
			"message": "Hello Vue!",
		},
		"methods": map[string]interface{}{
			"reverseMessage": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				this.Set("message", reverse(this.Get("message").String()))
				return nil
			}),
		},
	}
	_ = vue.New(m)

	select {}
}