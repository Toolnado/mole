package transport

type Transport interface {
	ListenAndServe() error
}
