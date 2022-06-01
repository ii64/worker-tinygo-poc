//go:build js && wasm
// +build js,wasm

package httpx

import "syscall/js"

var (
	Uint8Array = js.Global().Get("Uint8Array")
	Object     = js.Global().Get("Object")
)

func MapKey(v js.Value) (ret []string) {
	listKeys := Object.Call("keys", v)
	for i := 0; i < listKeys.Length(); i++ {
		elem := listKeys.Index(i)
		if !elem.IsUndefined() {
			ret = append(ret, elem.String())
		}
	}
	return
}
