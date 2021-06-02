package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request){
	myTemplates.ExecuteTemplate(w, "upload_form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request){
	file, fileHeader, err := r.FormFile("file")
	if err !=nil{
		panic(err)
	}
	create, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		return
	}
	_, err = io.Copy(create, file)
	if err != nil {
		return
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name" : name,
		"File" : "/static/" + fileHeader.Filename,
	})
}
func TestUploadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static",http.FileServer(http.Dir("./resources"))))
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}

//go:embed resources/4e1b305ff822ee15780d8350ada775c8.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	writer.WriteField("name", "Argadhana")
	file, _ := writer.CreateFormFile("file", "contohupload.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}