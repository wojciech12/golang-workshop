package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestWordCount(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://my?s=wwwww%20www", nil)
	expected := 2

	recorder := httptest.NewRecorder()
	wc := WCService{}

	wc.handleWordCount(recorder, r)

	if recorder.Code != http.StatusOK {
		t.Fatal(recorder.Code, recorder.Body.String())
	}

	ret := recorder.Body.String()
	num, err := strconv.Atoi(ret)

	if err != nil {
		t.Fatalf("Cannot parse the result as a int %s", ret)
	}

	if num != expected {
		t.Fatalf("Expected %d but got %d", expected, num)
	}
}
