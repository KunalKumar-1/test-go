package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/goodbye", handleGoodbye)

	fmt.Println("Listening on port 4000")

	log.Fatal(http.ListenAndServe(":4000", mux))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested path:", r.URL.Path)

	wc, err := w.Write([]byte("Hello world is served at health\n"))
	if err != nil {
		slog.Error("Error serving the health_handler err: " + err.Error())
		return
	}

	fmt.Println(wc, "Bytes written successfully")
}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested path:", r.URL.Path)

	wc, err := w.Write([]byte("Goodbye world is served at goodbye\n"))
	if err != nil {
		log.Fatal("Error serving the goodbye handler err: " + err.Error())
		return
	}
	fmt.Println(wc, "Bytes written successfully")
}
