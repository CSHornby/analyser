package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestUploadHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/home.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store}
	handler.Process(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Home page \n", string(w.Body.Bytes()))
}
