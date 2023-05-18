package main

import (
	"TroWeb/troWeb"
)

func main() {
	e := troWeb.New()
	e.GET("/", func(c *troWeb.Context) {
		c.String(200, "hello world\n")
	})
	e.GET("/hello", func(c *troWeb.Context) {
		c.Data(200, []byte("hi"))
	})
	e.Run(":9999")
}
