package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	handleRoot(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Welcome to our HomePage!\n")
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

func TestHandleHelloParameterized(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello?user=Testing", nil)

	handleHelloParameterized(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code:  expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello Testing!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleHelloNoParameterized(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello/", nil)

	handleHelloParameterized(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code:  expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello User!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleHelloWrongParameterized(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil)

	handleHelloParameterized(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code:  expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello User!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}
