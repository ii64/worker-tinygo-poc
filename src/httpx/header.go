package httpx

type Header map[string][]string

func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

func (h Header) Get(key, value string) string {
	values, ok := h[key]
	if !ok {
		return ""
	}
	if len(values) < 0 {
		return ""
	}
	return values[0]
}
