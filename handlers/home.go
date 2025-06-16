package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

type HomeHandler struct {
	Session *sessions.CookieStore
	Temp    *template.Template
}

func (h HomeHandler) Process(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session.Get(r, "analyse-session")
	// Check for flash messages
	flash := ""
	if flashes := sess.Flashes("error"); len(flashes) > 0 {
		flash = flashes[0].(string)
		sess.Save(r, w)
	}
	err := h.Temp.Execute(w, flash)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
