//go:build js && wasm
// +build js,wasm

package httpx

import (
	"strings"
	"syscall/js"
)

func (h Header) applyTo(jsObj js.Value) {
	// TODO: multi value header
	for k, v := range h {
		if len(v) < 1 {
			continue
		}
		jsObj.Set(k, v[0])
	}
}

type directJSHeader struct {
	jsObj js.Value
}

func (s directJSHeader) Set(key, value string) {
	s.jsObj.Call("set", js.ValueOf(key), js.ValueOf(value))
}

func (s directJSHeader) Add(key, value string) {
	s.jsObj.Call("append", js.ValueOf(key), js.ValueOf(value))
}

func (s directJSHeader) Del(key string) {
	s.jsObj.Call("delete", js.ValueOf(key))
}

func (s directJSHeader) Get(key string) string {
	val := s.jsObj.Call("get", js.ValueOf(strings.ToLower(key)))
	if !val.IsUndefined() {
		return val.String()
	}
	return ""
}
