package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request){
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/header", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request){
	w.Header().Add("author-by", "Argadhana")
	fmt.Fprintf(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/header", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()

	fmt.Println(response.Header.Get("author-by"))

}