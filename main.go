package main

import (
	"github.com/TRO148/troWeb/troWeb"
	"net/http"
)

func main() {
	r := troWeb.New()
	r.GET("/index", func(c *troWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *troWeb.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *troWeb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *troWeb.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *troWeb.Context) {
			c.JSON(http.StatusOK, troWeb.J{
				"username": c.PostForm("username"),
			})
		})

	}

	r.Run(":9999")
}
