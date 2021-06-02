package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerRequest(w http.ResponseWriter,  r *http.Request){
	fmt.Fprint(w, "HELLOWORLD")
}

func TestHttpRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hi", nil)
	recorder := httptest.NewRecorder()

	handlerRequest(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}