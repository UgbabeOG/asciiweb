package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/UgbabeOG/asciiweb/internals/renderer"
)


type PageData struct {
	ErrorMsg    string
	ASCIIart    string
	InputText   string
	BannerStyle string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "bad request", http.StatusMethodNotAllowed)
		return
	}
	if err := tmpl.Execute(w, PageData{}); err != nil {
		http.Error(w, "unable to serve", http.StatusInternalServerError)
	}
}

func HandleASCII(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}
	input := r.FormValue("text")
	banner := r.FormValue("banner")
	data := PageData{}
	ASCIIart, err := renderer.GenerateASCII(input, banner)
	if err != nil {

		data.ErrorMsg = err.Error()
		http.Error(w, "Banner File not found", http.StatusNotFound)
		return
	}
	data.ASCIIart = ASCIIart
	data.InputText = input
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "unable to execute", http.StatusInternalServerError)
	}
}

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "unable to parse form.", http.StatusBadRequest)
		return
	}

	input := r.FormValue("text")
	banner := r.FormValue("banner")
	data := PageData{
		InputText:   input,
		BannerStyle: banner,
	}
	ASCII, err := renderer.GenerateASCII(data.InputText, data.BannerStyle)
	if err != nil {
		data.ErrorMsg = err.Error()
		http.Error(w, data.ErrorMsg, http.StatusBadRequest)
		return
	}
	data.ASCIIart = ASCII
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data.ASCIIart)))
	w.Write([]byte(data.ASCIIart))
}
