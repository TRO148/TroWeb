package troWeb

import (
	"net/http"
	"path"
)

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

// Static 静态文件处理，参数为网址相对路径与服务器绝对路由
func (group *routerGroup) Static(relativePath string, root string) {
	//创建静态文件处理器，将服务器绝对路由与网址相对路径绑定生成处理函数
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	//注册GET请求
	group.GET(urlPattern, handler)
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

// createStaticHandler 创建静态文件处理器
// 解析请求地址，映射到服务器文件的真实地址
func (group *routerGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	//使地址与文件处理handler匹配
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		file := context.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			context.Status(http.StatusNotFound)
			return
		}
		//启动文件处理器
		fileServer.ServeHTTP(context.Writer, context.Request)
	}
}

// addRoute 添加路由
func (group *routerGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	//存入到路由树中engine一样，路由树一样
	group.engine.r.addRoute(method, pattern, handler)
}
