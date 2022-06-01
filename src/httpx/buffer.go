//go:build js && wasm
// +build js,wasm

package httpx

import (
	"bytes"
	"syscall/js"
)

type Buffer struct {
	bytes.Buffer
	closerFunc func(data js.Value) error
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Close() error {
	if f := b.closerFunc; f != nil {
		var data = Uint8Array.New(b.Len())
		js.CopyBytesToJS(data, b.Bytes())
		return f(data)
	}
	return nil
}
