const $ = require("shelljs")

$.mkdir('-p', 'build')

// See more at https://tinygo.org/docs/reference/usage/important-options/

const args = [
    "tinygo build",
        // "-size", "full",
        "-no-debug",
        "-gc", "leaking",
        "-opt", "z",
        "-o", "build/main.wasm",
        "-target", "wasm",
        "template/src"
]
$.exec(args.join(" "))

// $.exec(`sh -c "GOOS=js GOARCH=wasm go build -o build/main.wasm template/src"`)
$.exec("wasm-opt -O build/main.wasm -o build/main.wasm")
