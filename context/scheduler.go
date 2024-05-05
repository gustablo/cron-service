package context

type (
	Scheduler interface {
		InsertConcurrently(newJob interface{}) // need to find a way to do not use interface{}
	}
)
