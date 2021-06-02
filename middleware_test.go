package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type logMiddleware struct {
	Handler http.Handler
}

type errorHandler struct {
	Handler http.Handler
}

func (errorHandler *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err !=nil{
			fmt.Println("terjadi error!")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

	func (middleware logMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println("Before Middleware")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Middleware")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprintf(writer, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Execute")
		panic("Ups Error!")
	})

	logMiddleware := &logMiddleware{
		Handler: mux,
	}

	errorHandler := &errorHandler{
		Handler:logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		return 
	}
	
}