package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type requestData struct {
	Text string `json:"input"`
}

type responseData struct {
	Length int `json:"length"`
}

func stringLengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Store request data in requestData variable.
	var reqData requestData
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// String length calculation
	length := len(reqData.Text)
	respData := responseData{Length: length}

	// Response will be json.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respData)
}

func main() {
	http.HandleFunc("/", stringLengthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
