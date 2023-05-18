package troWeb

import (
	"net/http"
)

// HandlerFunc 处理函数 用于处理http请求
type HandlerFunc func(http.ResponseWriter, *http.Request)

// New 构造Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// New 构造Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
