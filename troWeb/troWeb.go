package troWeb

import (
	"net/http"
)

// HandlerFunc 处理函数 用于处理http请求
type HandlerFunc func(*Context)

// New 构造Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// newContext 构造Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}

// newRouter 构造router
func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
