package main

import (
	"TroWeb/troWeb"
	"fmt"
	"net/http"
)

func main() {
	e := troWeb.New()
	e.GET("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", request.URL.Path)
	})
	e.Run(":9999")
}
