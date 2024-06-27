package main

import (
	"fmt"
	"miniweb"
	"net/http"
)

func main() {
	m := miniweb.New()

	m.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	})

	m.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	m.Run(":9999")
}
