package httpx

import "syscall/js"

func (r *Request) ApplyFrom(v js.Value) error {
	r.header = directJSHeader{v.Get("headers")}
	return nil
}
