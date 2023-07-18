package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Qtest)

	route := negroni.Classic()
	route.UseHandler(mux)
	route.Run(":3005")
}

func Qtest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Qtest!")
}
