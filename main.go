package main

import (
	"html/template"
	"log"
	"main/handlers"
	"main/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	store := sessions.NewCookieStore([]byte("never-store-the-key-like-this-:)"))
	temp, err := template.ParseFiles("templates/layout.html", "templates/upload.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	uploadHandler := handlers.UploadHandler{Session: store, Temp: temp, CsvService: services.ExtractCsv{}, Analyser: services.Analyse{}, Clean: services.Clean{}}

	temp, err = template.ParseFiles("templates/layout.html", "templates/home.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	homeHandler := handlers.HomeHandler{Session: store, Temp: temp}

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Process).Methods("GET")
	r.HandleFunc("/upload", uploadHandler.Process).Methods("POST")
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
