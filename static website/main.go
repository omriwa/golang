package main

import (
	// "fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return
		}
		http.ServeFile(w,r,"./static/index.html")
	})

	http.HandleFunc("/samp", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "samp")
	})

	log.Fatalln(http.ListenAndServe("localhost:3000", nil), "faild to setup server")
}
