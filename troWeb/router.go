package troWeb

// Router 路由，用于存储路由信息
type Router struct {
	handlers map[string]HandlerFunc
}

// addRoute 添加路由，将请求方法+请求路径作为key，处理函数作为value，存入map
func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handler
}

// handle 处理请求，根据请求方法+请求路径，从map中取出对应的处理函数，执行
func (router *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.handlers[key]; ok {
		handler(c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
