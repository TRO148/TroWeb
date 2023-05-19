package troWeb

// routerGroup 分组，提供引擎的基本方法
type routerGroup struct {
	prefix      string        //前缀
	middlewares []HandlerFunc //中间件，用于特殊处理
	parent      *routerGroup  //父组
	engine      *Engine       //所有分组使用同一个引擎
}

// Use 添加中间件
func (group *routerGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// Group 添加新的分组
// 所有分组使用同一个引擎
func (group *routerGroup) Group(prefix string) *routerGroup {
	engine := group.engine
	newGroup := &routerGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// addRoute 添加路由
func (group *routerGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	//存入到路由树中engine一样，路由树一样
	group.engine.r.addRoute(method, pattern, handler)
}

// GET 定义GET请求，查询数据
func (group *routerGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 定义POST请求，新建一个资源
func (group *routerGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// PUT 定义PUT请求，更新一个资源
func (group *routerGroup) PUT(pattern string, handler HandlerFunc) {
	group.addRoute("PUT", pattern, handler)
}

// PATCH 定义PATCH请求，更新一个资源的部分信息
func (group *routerGroup) PATCH(pattern string, handler HandlerFunc) {
	group.addRoute("PATCH", pattern, handler)
}
