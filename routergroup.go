package folcon

type IRouter interface {
	Use(middlewares ...HandlerFunc)
	Group(prefix string) *RouterGroup
	GET(path string, handlers ...HandlerFunc)
	POST(path string, handlers ...HandlerFunc)
}

type RouterGroup struct {
	Handlers HandlersChain
	BasePath string
	engine   *Engine
	parent   *RouterGroup
}

var _ IRouter = &RouterGroup{}

// Use adds middlewares to the router group
func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.Handlers = append(g.Handlers, middlewares...)
}

// Group creates a new sub-router group with a shared prefix
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		BasePath: g.BasePath + prefix,
		engine:   g.engine,
		parent:   g,
		Handlers: append(HandlersChain{}, g.Handlers...),
	}
}

// Internal method to handle route registration
func (g *RouterGroup) handle(httpMethod string, relativePath string, handlers HandlersChain) {
	absolutePath := g.BasePath + relativePath
	// Combine group-level middlewares with route-specific handlers
	mergedHandlers := append(g.Handlers, handlers...)
	g.engine.addRoute(httpMethod, absolutePath, mergedHandlers)
}

// GET is a shortcut for registering GET requests
func (g *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	g.handle("GET", path, handlers)
}

// POST is a shortcut for registering POST requests
func (g *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	g.handle("POST", path, handlers)
}
