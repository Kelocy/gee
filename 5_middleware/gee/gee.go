package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandleFunc // support middleware
		parent      *RouterGroup // support nesting, point to engine
		engine      *Engine      // all groups share a Engine instance
	}

	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// Constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup.
// All groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	// Get middlewares based on prefix
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
