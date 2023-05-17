package troWeb

import (
	"fmt"
	"net/http"
)

// HandlerFunc 处理函数 用于处理http请求
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由，将请求方法+请求路径作为key，处理函数作为value，存入map
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 定义GET请求，查询数据
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义POST请求，新建一个资源
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// PUT 定义PUT请求，更新一个资源
func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute("PUT", pattern, handler)
}

// PATCH 定义PATCH请求，更新一个资源的部分信息
func (engine *Engine) PATCH(pattern string, handler HandlerFunc) {
	engine.addRoute("PATCH", pattern, handler)
}

// DELETE 定义DELETE请求，删除一个资源
func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler)
}

// Run 启动http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 用于ListenAndServe调用，实现ServeHTTP接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}