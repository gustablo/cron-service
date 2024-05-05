package context

type Env interface {
	Get(path string) string
}
