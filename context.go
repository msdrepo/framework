package folcon

import (
	"net/http"
)

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

type Context struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	handlers HandlersChain
	index    int
}

// Next allows running the next middleware in the chain
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

func (c *Context) String(code int, format string, values ...any) {
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(format))
}
