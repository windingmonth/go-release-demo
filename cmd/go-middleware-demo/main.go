package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.HandlerFunc(MyMiddleware1))
	n.Use(negroni.HandlerFunc(MyMiddleware2))
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":3004",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	// n.Run(":3004")
	// http.ListenAndServe(":3000", n)
}
func MyMiddleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Print("MyMiddleware before")
	next(rw, r)
	fmt.Print("MyMiddleware after")
}
func MyMiddleware2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Print("MyMiddleware before")
	next(rw, r)
	fmt.Print("MyMiddleware after")
}
