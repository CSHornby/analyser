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
	// Check for flash messages
	sess, _ := h.Session.Get(r, "analyse-session")
	flash := ""
	if flashes := sess.Flashes("error"); len(flashes) > 0 {
		flash = flashes[0].(string)
		sess.Save(r, w)
	}
	h.Temp.Execute(w, flash)
}
