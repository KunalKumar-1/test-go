package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleHelloParameterized)

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

	var output bytes.Buffer
	output.WriteString("Hello ")
	output.WriteString(username)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error starting response body", "err", err.Error())
		return
	}
}
