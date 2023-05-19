package main

import (
	"github.com/TRO148/troWeb/troWeb"
	"log"
	"net/http"
	"time"
)

func onlyForV2() troWeb.HandlerFunc {
	return func(c *troWeb.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := troWeb.New()
	r.Use(troWeb.Logger()) // global midlleware
	r.GET("/", func(c *troWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *troWeb.Context) {
			// expect /hello/troWebktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
