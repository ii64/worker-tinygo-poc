import "./wasm_exec.js"

let program = () => {
	return async (request) => {
		const go = new Go()
		const ins = await WebAssembly.instantiate(wasm, go.importObject)
		go._inst = ins
		go._values = [ // JS values that Go currently has references to, indexed by reference id
					NaN,     // 0
					0,       // 1
					null,    // 2
					true,    // 3
					false,   // 4
					global,  // 5
					go,      // 6 - jsGo

					// custom
					request,  // 7
				];
		go._goRefCounts = []; // number of references that Go has to a JS value, indexed by reference id
		go._ids = new Map();  // mapping from JS values to reference ids
		go._idPool = [];      // unused ids that have been garbage collected
		go.exited = false;    // whether the Go program has exited

		// execute handler
		go._inst.exports._start();
	}
}


addEventListener('fetch', (event) => {
	event.respondWith(handleRequest(event))
})



async function handleRequest(event) {

	let request = event.request
	let response = {
		status: 404,
		headers: {}
	};
	request.response = response

	const handler = program()
	await handler(request)

	const { body, ...opt } = response
	return new Response(body || "unknown handler", opt)

}
