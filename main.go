package main

import (
	"github.com/TRO148/troWeb/troWeb"
	"log"
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
	r.Static("/", "D:/Learn/TroWeb")

	r.Run(":9999")
}
