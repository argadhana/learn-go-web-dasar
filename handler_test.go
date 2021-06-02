package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handlefunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handlefunc,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}