package troWeb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context 上下文，包装原生的http.ResponseWriter和*http.Request
type Context struct {
	// 原生的参数
	Writer  http.ResponseWriter
	Request *http.Request

	// 请求信息
	Path   string
	Method string
	Params map[string]string

	// 响应信息
	StatusCode int

	//中间件
	handlers []HandlerFunc
	// index 用于记录当前执行到第几个中间件
	index int

	// engine指针
	engine *Engine
}

// Fail 报错
func (c *Context) Fail(statusCode int, message string) {
	c.index = len(c.handlers) - 1
	c.JSON(statusCode, J{"message": message})
}

// Next 向下执行中间件
// 当中间件调用Next()的时候，向后执行，然后再回来
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// Param 获取路由参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 获取表单参数
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// Query 获取查询参数
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Status 设置响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 多种响应体写入方式

// String 设置响应体
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 设置响应体
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// 将obj序列化为json格式
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 设置响应体
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 设置响应体,name为载入模板的项目名称，data为传入模板的数据
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	// 执行模板
	err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data)
	if err != nil {
		c.Fail(500, err.Error())
	}
}
