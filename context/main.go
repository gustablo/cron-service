package context

import (
	"database/sql"
	"sync"
)

type Context struct {
	Scheduler Scheduler
	DB        *sql.DB
	Env       Env
}

var (
	context *Context
	mutex   sync.Mutex
	once    sync.Once
)

func NewContext(s Scheduler, db *sql.DB, env Env) *Context {
	mutex.Lock()
	defer mutex.Unlock()
	once.Do(func() {
		context = &Context{
			Scheduler: s,
			DB:        db,
			Env:       env,
		}
	})

	return context
}

func GetContext() *Context {
	return context
}
