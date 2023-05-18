package troWeb

import (
	"net/http"
	"strings"
)

// Router 路由，用于存储路由信息
type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// parsePattern 解析路由，辅助函数，将pattern拆分parts
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	//所以不支持/*hi/hihi/hi,没有后续，只能到/*hi
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

// addRoute 添加路由，存入路由树与handlers方法中
func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	//存入节点，以方法为树的基本节点
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)

	//存入处理函数
	key := method + "-" + pattern
	router.handlers[key] = handler
}

// getRoute 获取路由，查询节点与参数
func (router *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}

	//查询到这个点
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			//如果为:，则添加到参数
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				//如果为*，把路径后续添加到参数
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// handle 处理请求，根据请求方法+请求路径，从map中取出对应的处理函数，执行
func (router *Router) handle(c *Context) {
	n, params := router.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		router.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
