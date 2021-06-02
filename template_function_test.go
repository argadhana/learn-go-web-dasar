package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (mypage MyPage) SayHello(name string) string  {
	return "Hello " + name + ", My name is " + mypage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Arga"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Affan",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//GlobalFunction
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.New("FUNCTION").Parse(`{{len "Arga"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Affan",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//Make Global Function
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request)  {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper" : func(value string) string{
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arga",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//Function Pipeline
func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request)  {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello" : func(value string) string{
			return "Hello " + value
		},
	})
	t = t.Funcs(map[string]interface{}{
		"upper" : func(value string) string{
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Arga",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}