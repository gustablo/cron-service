package context

type Server interface {
	ServeHTTP()
}
