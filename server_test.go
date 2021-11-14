package main

import (
	"fmt"
	"io/ioutil"
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
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("save is not 200")
	}

	key := keyBuilder.Get()
	savedMessage, _ := keeper.Get(key)
	if savedMessage != testMessage {
		t.Error("message was not saved")
	}
	result := w.Result()
	defer result.Body.Close()
	data, _ := ioutil.ReadAll(result.Body)
	if !strings.Contains(string(data), key) {
		t.Error("error: page without key")
	}
}
