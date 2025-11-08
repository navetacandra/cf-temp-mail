package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const workerKey = ""

func main() {
	http.HandleFunc("/mail-worker-webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Body == nil || r.Header.Get("X-Worker-Id") != workerKey {
			w.WriteHeader(418) // send teapot status code
			return
		}

		w.WriteHeader(200)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()
		fmt.Printf("%s\n", string(body))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
