package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type UserData struct {
	Name string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleHelloParameterized)
	mux.HandleFunc("/responses/{user}/hello/", handleUserResponsesHello)
	mux.HandleFunc("/user/hello", handleHelloHeader)
	mux.HandleFunc("/json", handleJSON)

	fmt.Println("Listening on port 4000")

	log.Fatal(http.ListenAndServe(":4000", mux))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested path:", r.URL.Path)

	_, err := w.Write([]byte("Welcome to our HomePage!\n"))
	if err != nil {
		slog.Error("Error serving the health_handler err: " + err.Error())
		return
	}

}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested path:", r.URL.Path)

	_, err := w.Write([]byte("Goodbye world is served at goodbye\n"))
	if err != nil {
		log.Fatal("Error serving the goodbye handler err: " + err.Error())
		return
	}
}

func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested Path: ", r.URL.Path)

	params := r.URL.Query()
	userlist := params["user"]

	username := "User"
	if len(userlist) > 0 {
		username = userlist[0]
	}

	handleHello(w, username)
}

func handleUserResponsesHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested Path: ", r.URL.Path)

	username := r.PathValue("user")

	handleHello(w, username)
}

func handleHelloHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested Path: ", r.URL.Path)
	//username := r.PathValue("user")
	username := r.Header.Get("user")
	if username == "" {
		http.Error(w, "invalid username provided", http.StatusBadRequest)
		return
	}

	handleHello(w, username)
}

func handleHelloNoHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested Path: ", r.URL.Path)
	//username := r.PathValue("user")
	username := r.Header.Get("user")
	if username == "" {
		http.Error(w, "invalid username provided", http.StatusBadRequest)
		return
	}

	handleHello(w, username)
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested Path: ", r.URL.Path)

	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading request body", "err: ", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	if len(byteData) == 0 {
		http.Error(w, "empty request body", http.StatusBadRequest)
		return
	}

	var reqData UserData
	err = json.Unmarshal(byteData, &reqData)
	if err != nil {
		slog.Error("error unmarshalling request body", "err", err)
		http.Error(w, "error parsing request body", http.StatusBadRequest)
		return
	}

	if reqData.Name == "" {
		http.Error(w, "invalid request body!", http.StatusBadRequest)
		return
	}

	handleHello(w, reqData.Name)

}

func handleHello(w http.ResponseWriter, username string) {

	var output bytes.Buffer
	output.WriteString("Hello ")
	output.WriteString(username)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error starting response body", "err: ", err.Error())
		return
	}
}
