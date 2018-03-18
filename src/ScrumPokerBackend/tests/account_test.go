package tests

import (
	"ScrumPokerBackend/app"
	"ScrumPokerBackend/routers"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a app.App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	//userJson := `{"username": "dennis", "balance": 200}`

	//reader := strings.NewReader(userJson) //Convert string to reader

	req, _ := http.NewRequest("GET", "/account", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func init() {

	router := mux.NewRouter()
	routers.AddAccountRouters(router)
	routers.AddUserRouters(router)
	httptest.NewServer(router)

}
