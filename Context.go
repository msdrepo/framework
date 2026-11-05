package folcon

import (
	"encoding/json"
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

func (c *Context) String(code int, format string, values ...any) {
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(format))
}

// JSON writes any data structure as a JSON response
func (c *Context) JSON(code int, data interface{}) {

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	c.Writer.WriteHeader(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// Next allows running the next middleware in the chain
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// BindJSON parses the request body JSON into the given struct pointer
func (c *Context) BindJSON(obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	return decoder.Decode(obj)
}

// Query gets URL query parameters (e.g. ?name=augustin)
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}
