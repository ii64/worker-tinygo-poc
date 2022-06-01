//go:build js && wasm
// +build js,wasm

package httpx

import "syscall/js"

func (h Header) applyTo(jsObj js.Value) {
	for k, v := range h {
		if len(v) < 1 {
			continue
		}
		jsObj.Set(k, v[0])
	}
}
