package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestManageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:7700/manager", nil)
	if err != nil {
		t.Fatalf("could not create the request: %v", err)
	}
	rec := httptest.NewRecorder()

	managerHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

}
