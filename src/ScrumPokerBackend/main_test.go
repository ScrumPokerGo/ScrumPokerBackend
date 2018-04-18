package main_test

import (
	"os"
	"testing"
	"ScrumPokerBackend/app"
	"net/http"
	"net/http/httptest"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func TestProjectList(t *testing.T) {
	req, _ := http.NewRequest("GET", "/project", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}