package troWeb

import (
	"log"
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
