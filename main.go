package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// Request struct to hold the incoming JSON request
type Request struct {
	Command string `json:"command"`
}

// Response struct to hold the response JSON
type Response struct {
	Output        string `json:"output"`
	Error         string `json:"error,omitempty"`
	ExecutionTime string `json:"executionTime"`
}

func executeCommand(w http.ResponseWriter, r *http.Request) {
	var req Request

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := exec.CommandContext(r.Context(), "bash", "-c", req.Command)
	log.Printf("Executing command: %s\n", req.Command)
	// Execute the command
	startTime := time.Now()
	out, err := cmd.CombinedOutput()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	response := Response{
		Output:        string(out),
		ExecutionTime: executionTime.String(),
	}

	// If there was an error, add it to the response
	if err != nil {
		response.Error = err.Error()
	}

	log.Printf("Execution completed (time: %s): %s\n", executionTime, req.Command)
	// Convert response to JSON
	respJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func main() {
	http.HandleFunc("/execute", executeCommand)

	log.Println("Server starting on port 12084...")
	if err := http.ListenAndServe(":12084", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
