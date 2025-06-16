package handlers

import (
	"html/template"
	"main/services"
	"net/http"

	"github.com/gorilla/sessions"
)

type UploadHandler struct {
	Session    *sessions.CookieStore
	Temp       *template.Template
	CsvService services.ExtractCsvI
	Analyser   services.AnalyseI
	Clean      services.CleanI
}

func (u UploadHandler) Process(w http.ResponseWriter, r *http.Request) {
	var records [][]string
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	switch fileHeader.Header.Get("Content-Type") {
	case "text/csv":
		records, err = u.CsvService.Extract(file)
		if err != nil {
			flashAndRedirect(u, w, r)
			return
		}
	default:
		flashAndRedirect(u, w, r)
		return
	}

	cleaned, err := u.Clean.Clean(records)
	if err != nil {
		flashAndRedirect(u, w, r)
		return
	}

	categories := u.Analyser.Analyse(cleaned)

	u.Temp.Execute(w, categories)
}

func flashAndRedirect(u UploadHandler, w http.ResponseWriter, r *http.Request) {
	sess, _ := u.Session.Get(r, "analyse-session")
	sess.AddFlash("Unsupported file type. Please upload a CSV file.", "error")
	sess.Save(r, w) // Save session after adding flash message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
