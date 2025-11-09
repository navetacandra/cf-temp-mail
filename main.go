package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type WebhookData struct {
	Id      string `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

var clients map[string]chan string = make(map[string]chan string)

func main() {
	workerKey, exists := os.LookupEnv("WORKER_KEY")
	if !exists {
		fmt.Println("Worker key not set!")
	}

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

		var tmp WebhookData
		if err := json.Unmarshal(body, &tmp); err != nil {
			return
		}

		at := strings.IndexByte(tmp.To, '@')
		if at == -1 {
			return
		}

		recipient := tmp.To[:at]
		_, exist := clients[recipient]
		if !exist {
			return
		}

		clients[recipient] <- string(body)
	})

	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Cache-Control", "no-cache")

		subscribeFor := r.URL.Query().Get("for")
		if r.Method != "GET" || len(subscribeFor) < 1 {
			w.WriteHeader(418) // send teapot status code
			return
		}

		rc := http.NewResponseController(w)
		requestDone := r.Context().Done()

		message, exist := clients[subscribeFor]
		if !exist {
			clients[subscribeFor] = make(chan string)
			message = clients[subscribeFor]
		}

		// intial ping
		_, err := fmt.Fprintf(w, "ping: 1\n\n")
		if err != nil {
			return
		}
		err = rc.Flush()
		if err != nil {
			return
		}

		for {
			select {
			case <-requestDone:
				delete(clients, subscribeFor)
				return
			case msg := <-message:
				_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
				if err != nil {
					return
				}
				err = rc.Flush()
				if err != nil {
					return
				}
			}
		}

	})

	fileHandler := http.FileServer(http.Dir("web"))
	http.Handle("/", fileHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
