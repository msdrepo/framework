package folcon

import (
	"net/http"
)

const Version = "0.1.0"

type Engine struct {
	*RouterGroup
	routes map[string]HandlersChain
}

// New creates a new instance of the Folcon engine
func New() *Engine {
	engine := &Engine{
		routes: make(map[string]HandlersChain),
	}
	engine.RouterGroup = &RouterGroup{
		BasePath: "",
		engine:   engine,
	}
	return engine
}

// Registers the route in the engine's map
func (engine *Engine) addRoute(method string, path string, handlers HandlersChain) {
	key := method + ":" + path
	engine.routes[key] = handlers
}

// ServeHTTP satisfies the http.Handler interface
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + ":" + req.URL.Path
	if handlers, exists := engine.routes[key]; exists {
		c := &Context{
			Writer:   w,
			Request:  req,
			handlers: handlers,
			index:    -1,
		}
		c.Next()
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Page Not Found"))
	}
}

// Run starts the HTTP server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
