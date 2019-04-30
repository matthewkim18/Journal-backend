package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestJournalIndex(t *testing.T) {
	testJournal := &Journal{
		ID:   "111",
		Subject:  "test title",
		Text: "test author",
	}
	journal_list["111"] = testJournal
	req1, err := http.NewRequest("GET", "/journals", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "GET", "/Journals", FindJournal)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}
	// expected response
	er1 := "{\"meta\":null,\"data\":[{\"id\":\"111\",\"subject\":\"test title\",\"text\":\"test author\"}]}\n"
	if rr1.Body.String() != er1 {
		t.Error("Response body does not match")
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}