package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsForMonthHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(eventsForMonthHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/events_for_month/e2")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
}
