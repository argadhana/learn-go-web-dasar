package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)
type Action struct {
	Title string
	Name string
}
func TemplateActionIf(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Action{
		Title: "Template data struct",
		Name:"Argadhana",
	})
}
func TestTemplataActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionComparator(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title": "Template action operator",
		"FinalValue": 90,
	})
}

func TestTemplataActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml",map[string]interface{}{
		"Title": "Range action template",
		"Hobbies": []string{
			"Traveling", "Fotografi", "Web Design",
		},
	})
}

func TestTemplataActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Data struct{
	Title string
	Name string
	Address Address
}

type Address struct {
	Street string
	City string
}
func TemplateActionWith(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", Data{
		Title: "Template action with",
		Name: "Argadhana",
		Address: Address{
			Street: "Jl.Mana aja si",
			City: "Semarang",
		},
	})
}

func TestTemplataActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}