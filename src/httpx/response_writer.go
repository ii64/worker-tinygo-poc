package httpx

type ResponseWriter interface {
	Write(b []byte) (int, error)
	WriteHeader(s int)
	Header() Header

	Close() error
}
