package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	handleHealth(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello world is served at health\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleGoodbye(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/goodbye", nil)
	handleGoodbye(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Goodbye world is served at goodbye\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}
