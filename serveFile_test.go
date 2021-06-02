package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}

// With Golang Embed

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServeFileWithEmbed(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileWithEmbed),
	}

	err := server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}