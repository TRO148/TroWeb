package troWeb

import (
	"net/http"
)

// J 便于写JSON
type J map[string]interface{}

// HandlerFunc 处理函数 用于处理http请求
type HandlerFunc func(*Context)

// New 构造Engine
func New() (engine *Engine) {
	engine = &Engine{r: newRouter()}
	engine.routerGroup = &routerGroup{engine: engine}
	engine.groups = []*routerGroup{engine.routerGroup}
	return engine
}

// newContext 构造Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
		index:   -1, //中间件
	}
}

// newRouter 构造router
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
