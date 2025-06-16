package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestHomeProcess(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/home.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := HomeHandler{Temp: temp, Session: store}
	handler.Process(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Home page \n", string(w.Body.Bytes()))
}

func TestHomeProcessFlash(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/home.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	sess, _ := store.Get(req, "analyse-session")
	sess.AddFlash("Errors", "error")
	handler := HomeHandler{Temp: temp, Session: store}
	handler.Process(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Home page Errors\n", string(w.Body.Bytes()))
}
