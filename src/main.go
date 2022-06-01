//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
	"template/src/httpx"
	"unsafe"
)

func main() {
	initHandler()

	req := predefObjValue(7)
	res := req.Get("response")

	method := req.Get("method").String()
	url := req.Get("url").String()

	{
		req, err := httpx.NewRequest(method, url, nil)
		if err != nil {
			panic("cannot parse request: " + err.Error())
		}
		resWriter := httpx.NewResponseWriter(res)
		Handler(resWriter, req)
	}
}

// predefObjValue
// https://cs.opensource.google/go/go/+/19309779ac5e2f5a2fd3cbb34421dafb2855ac21:src/syscall/js/js.go;l=95;bpv=1;bpt=0
func predefObjValue(id uint32) (val js.Value) {
	var base = js.Global()
	var refBase = *(*uint64)(unsafe.Pointer(&base))
	*(*uint64)(unsafe.Pointer(&val)) = ((refBase >> 2) << 2) | uint64(id)
	return
}
