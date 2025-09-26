package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Accept only POST
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Optional: limit body size
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB

	ct := r.Header.Get("Content-Type")
	switch {
	case ct == "application/json" || ct == "application/json; charset=utf-8":
		var p Person
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&p); err != nil {
			http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
			return
		}
		// respond with a confirmation
		resp := map[string]interface{}{
			"status":  "ok",
			"message": fmt.Sprintf("received JSON for %s (%d)", p.Name, p.Age),
			"when":    time.Now().UTC(),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)

	default:
		// Try to read raw body (for form encoded or other)
		body, _ := io.ReadAll(r.Body)
		resp := map[string]interface{}{
			"status":  "ok",
			"message": "received raw body",
			"body":    string(body),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func main() {
	http.HandleFunc("/submit", submitHandler)

	addr := ":8080"
	log.Printf("starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
