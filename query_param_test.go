package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=arga", nil)
	recorder := httptest.NewRecorder()

	sayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	stringBody := string(body)
	fmt.Println(stringBody)
}

func MultipleParameterQuery(w http.ResponseWriter, r *http.Request)  {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?first_name=arga&last_name=dhana", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	stringBody := string(body)
	fmt.Println(stringBody)
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hello?name=arga&name=dhana", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	stringBody := string(body)
	fmt.Println(stringBody)
}