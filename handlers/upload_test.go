package handlers

import (
	"errors"
	"html/template"
	"io"
	"main/models"
	"main/services/mocks"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestUploadHandler(t *testing.T) {
	mockAnalyse := mocks.NewMockAnalyseI(t)
	mockClean := mocks.NewMockCleanI(t)
	mockExtract := mocks.NewMockExtractCsvI(t)

	records := [][]string{}
	cleaned := []models.Entry{}
	categories := map[string]float64{
		"Food":      100.0,
		"Utilities": 200.0,
	}
	pr, pw := io.Pipe()
	multipartWriter := multipart.NewWriter(pw)

	go func() {
		defer multipartWriter.Close()

		// Add form field
		filePart, err := multipartWriter.CreateFormFile("file", "file.csv")
		assert.Nil(t, err)
		_, err = filePart.Write([]byte("File content"))

		assert.Nil(t, err)
	}()

	req, err := http.NewRequest(http.MethodPost, "/upload", pr)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	file, fileHeader, err := req.FormFile("file")
	fileHeader.Header.Set("Content-Type", "text/csv")

	assert.Nil(t, err)
	mockExtract.On("Extract", file).Once().Return(records, nil)
	mockClean.On("Clean", records).Once().Return(cleaned, nil)
	mockAnalyse.On("Analyse", cleaned).Once().Return(categories, nil)

	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/upload.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store, CsvService: mockExtract, Analyser: mockAnalyse, Clean: mockClean}
	handler.Process(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Food-100|Utilities-200|\n", string(w.Body.Bytes()))
}

func TestUploadHandlerNoFile(t *testing.T) {
	mockAnalyse := mocks.NewMockAnalyseI(t)
	mockClean := mocks.NewMockCleanI(t)
	mockExtract := mocks.NewMockExtractCsvI(t)

	req, err := http.NewRequest(http.MethodPost, "/upload", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/upload.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store, CsvService: mockExtract, Analyser: mockAnalyse, Clean: mockClean}
	handler.Process(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUploadHandlerWithBadFileType(t *testing.T) {
	mockAnalyse := mocks.NewMockAnalyseI(t)
	mockClean := mocks.NewMockCleanI(t)
	mockExtract := mocks.NewMockExtractCsvI(t)

	pr, pw := io.Pipe()
	multipartWriter := multipart.NewWriter(pw)

	go func() {
		defer multipartWriter.Close()

		// Add form field
		filePart, err := multipartWriter.CreateFormFile("file", "file.xls")
		assert.Nil(t, err)
		_, err = filePart.Write([]byte("File content"))

		assert.Nil(t, err)
	}()

	req, err := http.NewRequest(http.MethodPost, "/upload", pr)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	_, fileHeader, err := req.FormFile("file")
	fileHeader.Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/upload.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store, CsvService: mockExtract, Analyser: mockAnalyse, Clean: mockClean}
	handler.Process(w, req)
	assert.Equal(t, http.StatusSeeOther, w.Code)
	sess, _ := store.Get(req, "analyse-session")
	flashes := sess.Flashes("error")
	assert.Len(t, flashes, 1)
	assert.Equal(t, "Unsupported file type. Please upload a CSV file.", flashes[0])

}

func TestUploadHandlerBadCsv(t *testing.T) {
	mockAnalyse := mocks.NewMockAnalyseI(t)
	mockClean := mocks.NewMockCleanI(t)
	mockExtract := mocks.NewMockExtractCsvI(t)

	records := [][]string{}
	pr, pw := io.Pipe()
	multipartWriter := multipart.NewWriter(pw)

	go func() {
		defer multipartWriter.Close()

		// Add form field
		filePart, err := multipartWriter.CreateFormFile("file", "file.csv")
		assert.Nil(t, err)
		_, err = filePart.Write([]byte("File content"))

		assert.Nil(t, err)
	}()

	req, err := http.NewRequest(http.MethodPost, "/upload", pr)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	file, fileHeader, err := req.FormFile("file")
	fileHeader.Header.Set("Content-Type", "text/csv")

	assert.Nil(t, err)
	mockExtract.On("Extract", file).Once().Return(records, errors.New("Bad file"))

	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/upload.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store, CsvService: mockExtract, Analyser: mockAnalyse, Clean: mockClean}
	handler.Process(w, req)
	assert.Equal(t, http.StatusSeeOther, w.Code)
}

func TestUploadHandlerBadCsv2(t *testing.T) {
	mockAnalyse := mocks.NewMockAnalyseI(t)
	mockClean := mocks.NewMockCleanI(t)
	mockExtract := mocks.NewMockExtractCsvI(t)

	cleaned := []models.Entry{}

	records := [][]string{}
	pr, pw := io.Pipe()
	multipartWriter := multipart.NewWriter(pw)

	go func() {
		defer multipartWriter.Close()

		// Add form field
		filePart, err := multipartWriter.CreateFormFile("file", "file.csv")
		assert.Nil(t, err)
		_, err = filePart.Write([]byte("File content"))

		assert.Nil(t, err)
	}()

	req, err := http.NewRequest(http.MethodPost, "/upload", pr)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	file, fileHeader, err := req.FormFile("file")
	fileHeader.Header.Set("Content-Type", "text/csv")

	assert.Nil(t, err)
	mockExtract.On("Extract", file).Once().Return(records, nil)
	mockClean.On("Clean", records).Once().Return(cleaned, errors.New("Bad file"))

	w := httptest.NewRecorder()
	temp, err := template.ParseFiles("./../test-files/upload.html")
	store := sessions.NewCookieStore([]byte("test-key-test-key-test-key-test-key"))
	handler := UploadHandler{Temp: temp, Session: store, CsvService: mockExtract, Analyser: mockAnalyse, Clean: mockClean}
	handler.Process(w, req)
	assert.Equal(t, http.StatusSeeOther, w.Code)
}
