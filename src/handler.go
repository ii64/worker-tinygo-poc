package main

import (
	"bytes"
	"embed"
	"fmt"
	"runtime"
	"strconv"
	"template/src/httpx"

	qrcode "github.com/skip2/go-qrcode"
)

//go:embed tpx
var efs embed.FS

func initHandler() {
}

func pageIndex(w httpx.ResponseWriter, r *httpx.Request) {
	w.WriteHeader(200)
	// actually we can do with another way
	bs, err := efs.ReadFile("tpx/index.html")
	if err != nil {
		panic(err)
	}

	bs = bytes.Replace(bs, []byte("{{version}}"), []byte(runtime.Version()), -1)
	bs = bytes.Replace(bs, []byte("{{num_goroutine}}"), []byte(strconv.Itoa(runtime.NumGoroutine())), -1)

	bs = bytes.Replace(bs, []byte("{{goarch}}"), []byte(runtime.GOARCH), -1)
	bs = bytes.Replace(bs, []byte("{{goos}}"), []byte(runtime.GOOS), -1)

	w.Write(bs)
}

func pageGenQR(w httpx.ResponseWriter, r *httpx.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		w.WriteHeader(400)
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte("q is not specified"))
		return
	}

	w.WriteHeader(201)
	w.Header().Set("content-type", "image/png; charset=utf-8")

	bs, err := qrcode.Encode(q, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}
	w.Write(bs)
}

func Handler(w httpx.ResponseWriter, r *httpx.Request) {
	defer w.Close()

	fmt.Printf("Get request url = %s\n", r.URL.String())
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	switch r.URL.Path {
	case "/":
		pageIndex(w, r)
	case "/genqr":
		pageGenQR(w, r)
	}

}
