package handlers

import "net/http"

type HandlerI interface {
	Process(w http.ResponseWriter, r *http.Request)
}
