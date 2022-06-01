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
	jreq := predefObjValue(7)
	jres := jreq.Get("response")

	method := jreq.Get("method").String()
	url := jreq.Get("url").String()

	req, err := httpx.NewRequest(method, url, nil)
	if err != nil {
		panic("cannot parse request: " + err.Error())
	}
	if err = req.ApplyFrom(jreq); err != nil {
		panic("cannot apply from binding: " + err.Error())
	}
	resWriter := httpx.NewResponseWriter(jres)
	Handler(resWriter, req)
}

// predefObjValue
// https://cs.opensource.google/go/go/+/19309779ac5e2f5a2fd3cbb34421dafb2855ac21:src/syscall/js/js.go;l=95;bpv=1;bpt=0
func predefObjValue(id uint32) (val js.Value) {
	var base = js.Global()
	var refBase = *(*uint64)(unsafe.Pointer(&base))
	*(*uint64)(unsafe.Pointer(&val)) = ((refBase >> 2) << 2) | uint64(id)
	return
}
