package context

type Context struct {
	Scheduler Scheduler
	Server    Server
}

func CreateContext() *Context {
	return &Context{
		Scheduler: nil,
		Server:    nil,
	}
}

func (c *Context) Register(srv Server, scheduler Scheduler) {
	c.Scheduler = scheduler
	c.Server = srv
}
