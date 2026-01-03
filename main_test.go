package main

import (
	"bytes"
	"encoding/json"
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

func TestHandleUserResponsesHello(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/responses/TestMan/hello", nil)
	r.SetPathValue("user", "TestMan")
	w := httptest.NewRecorder()

	handleUserResponsesHello(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code:  expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello TestMan!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleHelloHeader(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/user/hello", nil)
	r.Header.Set("user", "TestMan")

	w := httptest.NewRecorder()

	handleHelloHeader(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello TestMan!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleHelloNoHeader(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/user/hello", nil)

	w := httptest.NewRecorder()

	handleHelloNoHeader(w, r)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %d, got %d\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("invalid username provided\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleJSON(t *testing.T) {
	testRequest := UserData{
		Name: "human",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatal("error marshalling testRequest data:", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/josn", bytes.NewBuffer(marshalledRequestBody))

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v got %v\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello human!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleJSONEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/json", nil)

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v got %v\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("empty request body\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}

func TestHandleJSONEmptyNameFeild(t *testing.T) {
	testRequest := UserData{
		Name: "",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatal("error marshalling testRequest data:", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(marshalledRequestBody))

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v got %v\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("invalid request body!\n")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response body: expected %s, got %s\nbody: %s\n",
			string(expectedMessage), string(w.Body.Bytes()), w.Body.String())
	}
}
