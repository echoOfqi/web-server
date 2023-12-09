package web

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandFunc)}
}
func (r *router) addRouter(method string, pattern string, handler HandFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
