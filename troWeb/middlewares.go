package troWeb

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// Logger 查看处理时间
func Logger() HandlerFunc {
	return func(c *Context) {
		//开始处理
		t := time.Now()
		c.Next()
		//结束处理
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

// Recovery 错误处理
func Recovery() HandlerFunc {
	return func(context *Context) {
		// 在任务执行过程中，如果出现panic，则会调用defer中的匿名函数
		defer func() {
			// 恢复panic
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				context.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		// 继续处理请求
		context.Next()
	}
}

// 打印错误
func trace(message string) string {
	var pcs [32]uintptr
	// Callers 用来返回调用栈的程序计数器
	n := runtime.Callers(3, pcs[:]) //跳过前三个caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")

	for _, pc := range pcs[:n] {
		// 根据程序计数器获取对应的函数
		fn := runtime.FuncForPC(pc)
		// 获取函数名
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
