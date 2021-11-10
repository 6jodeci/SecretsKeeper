package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleTestRequest(w *httptest.ResponseRecorder, r *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, r)
}

func TestIndexPageCase(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("index page is not 200")
	}
}

func TestSaveMessage(t *testing.T) {
	testMessage := "foo"
	postData := strings.NewReader(fmt.Sprintf("message=%s", testMessage))
	request, _ := http.NewRequest("POST", "/", postData)
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("save is not 200")
	}
}
