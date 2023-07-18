package main

import (
	"log"
	"net/http"
)

type DemoFunc struct {
	hello string
}

func (f DemoFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.hello))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", DemoFunc{hello: "你好"})
	mux.Handle("/hello/", DemoFunc{hello: "你好,anyway"})

	server := &http.Server{
		Addr:    ":8002",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
